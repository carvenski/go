# -*- coding: utf-8 -*-
import tornado
from tornado.gen import coroutine, multi


ioloop = tornado.ioloop.IOLoop()

def _stop(future):
    ioloop.stop()

def run_until_complete(future, ioloop=ioloop):
    """Keep running untill the future is done"""
    ioloop.add_future(future, _stop) 
    ioloop.start()

@coroutine
def producer(i):
    print('---------1', i)
    print('---------2', i)
    print('---------3', i)
    print('---------4', i)

@coroutine
def runner():
    print('--------start--------')
    for i in xrange(10):
        producer(i)  
    print('--------end--------')

run_until_complete(runner())

