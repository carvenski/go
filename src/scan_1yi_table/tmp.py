#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
https://movoto.atlassian.net/browse/DATA-1779
"""

from __future__ import unicode_literals
import sys

sys.path.append('../properties')
from import_pr_to_association_multi import process_one_row

from tars.aio.http import HTTPSessionClient 
from tars.aio import coroutine2, run_until_complete, Queue, SqlHelper, \
    sleep, return_in_coroutine, wait, set_aio_n_workers
from tars.utils import ThreadPoolExecutor
from tars.options import options as opt
from Utilities.movoto.logger import MLogger
import  Utilities as SS
from Utilities.movoto.settings import BACKEND_SERVICE_HOST

import time

@coroutine2
def scan_all_pr(db, task_q):
    start_id = opt.pr_start_id
    sql = '''
        SELECT pr.id, pr.property_id
        FROM public_record AS pr
        WHERE pr.id>'{pr_id}'
        ORDER BY pr.id ASC
        LIMIT {limit};
    '''
    try:
        while 1: 
            if task_q.full():
                print('-------------------------------------------------------------------full----')
            opt.logger.info('[scanning] pr data from start_id %s', start_id)
            _sql = sql.format(pr_id=start_id, limit=opt.limit)
            pr_list = yield db.getAllBySql(_sql, dict_result=True)
            if not pr_list:
                opt.logger.info('[done] scanning all pr data')
                break
            yield task_q.put(pr_list)
            start_id = pr_list[-1]["id"]
    except Exception as err:
        opt.logger.exception(err)


@coroutine2
def scan_all_property(db, task_q):
    start_id = opt.property_start_id
    sql = '''
        SELECT mp.id, mp.public_record_id 
        FROM mls_public_record_association AS mp
        WHERE mp.id>'{property_id}'
        ORDER BY mp.id ASC
        LIMIT {limit};
    '''
    try:
        while 1:
            opt.logger.info('[scanning] property data from start_id %s', start_id)
            _sql = sql.format(property_id=start_id, limit=opt.limit)
            property_list = yield db.getAllBySql(_sql, dict_result=True)

            if not property_list:
                opt.logger.info('[done] scanning all property data')
                break
            yield task_q.put(property_list)
            start_id = property_list[-1]["id"]
    except Exception as err:
        opt.logger.exception(err)


@coroutine2
def find_problem_pr(db, task_q, result_q):
    sql_to_find_if_p_exist = '''
        SELECT id FROM mls_public_record_association AS mp
        WHERE mp.id='{property_id}';
        '''
    sql_to_find_pr_multi_p = '''
        SELECT id FROM mls_public_record_association AS mp
        WHERE mp.public_record_id='{pr_id}';
        '''
    try:
        while not (opt.is_scan_all_pr_done and task_q.empty()):
            if task_q.empty():
                print('-----------------------------------------------empty----')
                yield sleep(1)
                continue

            batch = task_q.get_nowait()
            opt.logger.info('[checking] 10000 pr')
            t0 = time.time()
            for pr in batch:
                pr_id = pr["id"]
                pr_p_id = pr["property_id"]

                #1.pr have no property_id
                if not pr_p_id:
                    opt.logger.info("[found] pr:%s have no property_id ", pr_id)
                    continue
                #2.pr's property_id not exist in property table
                pr_p = yield db.getOneBySql(sql_to_find_if_p_exist.format(property_id=pr_p_id))
                if not pr_p:
                    opt.logger.info("[found] pr:%s 's property_id not exist ", pr_id)
                    continue
                #3.pr has multi properties
                pr_p_list = yield db.getAllBySql(sql_to_find_pr_multi_p.format(pr_id=pr_id), dict_result=True)
                if len(pr_p_list) > 1:
                    opt.logger.info("[found] pr:%s has multi properties ", pr_id)
                    continue
            t1 = time.time() 
            print('--------------------------------------------time: ', str(t1-t0))
    except Exception as err:
        opt.logger.exception(err)


@coroutine2
def find_problem_property(db, task_q, result_q):
    sql_to_find_if_pr_exist = '''
        SELECT id FROM public_record AS pr
        WHERE pr.id='{pr_id}';
        '''
    sql_to_find_pr_by_p_address = '''
        SELECT pr.id FROM mls_public_record_association AS mp
        INNER JOIN address AS ad1 ON mp.address_id=ad1.id
        INNER JOIN address AS ad2 ON ad1.address=ad2.address
        INNER JOIN public_record AS pr ON pr.address_id=ad2.id
        WHERE mp.id='{property_id}' ORDER BY pr.update_time DESC 
        LIMIT 1 ;
    '''
    sql_to_find_pr_p = '''
        SELECT property_id FROM public_record AS pr
        WHERE pr.id='{pr_id}';
        '''
    try:
        while not (opt.is_scan_all_property_done and task_q.empty()):
            t0 = time.time()
            if task_q.empty():
                yield sleep(1)
                continue

            batch = task_q.get_nowait()
            opt.logger.info('[checking] 10000 properties')
            for p in batch:
                p_id = p["id"]
                p_pr_id = p["public_record_id"]
                if not p_pr_id:
                    continue

                #4.property's pr_id not exist in pr table
                p_pr = yield db.getOneBySql(sql_to_find_if_pr_exist.format(pr_id=p_pr_id))
                if not p_pr:
                    opt.logger.info("[found] property:%s 's pr_id not exist", p_id)
                    pr = yield db.getOneBySql(sql_to_find_pr_by_p_address.format(property_id=p_id), dict_result=True)
                    if pr:
                        opt.logger.info("[found] pr:%s by property's address ", pr["id"])
                
                #5.property-A's pr_id exist in pr table but it point to Property-B 
                if p_pr:
                    # after scan pr table finished, pr_p here must be a normal property_id.
                    pr_p = yield db.getOneBySql(sql_to_find_pr_p.format(pr_id=p_pr_id))
                    if pr_p and pr_p[0] != p_id:
                        opt.logger.info("[found] p-A:%s has pr:%s has p-B:%s ", p_id, p_pr_id, pr_p[0])
            print('------------------------------------------------------------------time ', str(time.time()-t0))        
    except Exception as err:
        opt.logger.exception(err)


@coroutine2
def pr_runner(task_q, result_q, db):
    opt.set_option('logger', MLogger().getLogger('fix_data1779-scan_pr', ''))
    opt.logger.info("=====================scan pr table=======================")
    f1 = scan_all_pr(db, task_q)
    f2 = wait([find_problem_pr(db, task_q, result_q) for _ in range(50)])
    yield f1
    opt.is_scan_all_pr_done = True
    yield f2
    opt.is_find_problem_pr_done = True


@coroutine2
def property_runner(task_q, result_q, db):
    opt.set_option('logger', MLogger().getLogger('fix_data1779-scan_property', ''))
    opt.logger.info("=====================scan property table=======================")
    f1 = scan_all_property(db, task_q)
    f2 = wait([find_problem_property(db, task_q, result_q) for _ in range(50)])
    yield f1
    opt.is_scan_all_property_done = True
    yield f2
    opt.is_find_problem_pr_done = True


def main():
    task_q = Queue(100)
    result_q = None
    db = SqlHelper('movoto', use_connection_pool=True)
    set_aio_n_workers(n_workers=55)
    opt.set_option('pr_association_logger', MLogger().getLogger('fix_data1779-redo_pr_association', ''))

    opt.set_option('is_scan_all_pr_done', False)
    opt.set_option('is_scan_all_property_done', False)
    opt.set_option('is_find_problem_pr_done', False)
    
    opt.set_option('mode', None)
    opt.set_option('http_client', HTTPSessionClient())

    #1.scan pr table
    if opt.scan_pr:
        run_until_complete(pr_runner(task_q, result_q, db))
    
    #must reset is_find_problem_pr_done here
    opt.is_find_problem_pr_done = False
    
    #2.scan property table
    if opt.scan_property:
        run_until_complete(property_runner(task_q, result_q, db))


if __name__ == '__main__':
    try:
        opt.add_argument('--debug', required=False, default=False, action='store_true')
        opt.add_argument('--pr_start_id', type=str, default='')
        opt.add_argument('--property_start_id', type=str, default='')
        opt.add_argument('--limit', required=False, default='10000', type=str)
        opt.add_argument('--scan_pr', required=False, default=False, action='store_true')
        opt.add_argument('--scan_property', required=False, default=False, action='store_true')
        opt.parse_args()

        main()
    except Exception as e:
        opt.pr_association_logger.exception(e)
        sys.exit(2)

