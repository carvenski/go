package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
)

func main() {
	fmt.Println("[start] scan pr table")
	var q chan *[100][2]string = make(chan *[100][2]string, 1000)
	// db, err := sql.Open("mysql", "movoto:movoto123!@tcp(db3.ng.movoto.net)/movoto")
	db, err := sql.Open("mysql", "wentao-ro:admin@123@tcp(10.255.1.6)/movoto")
	db.SetMaxOpenConns(10) // use connection pool
	// db, err := sql.Open("mysql", "wentao-ro:admin@123@tcp(10.255.1.6)/movoto")
	// db, err := sql.Open("mysql", "root:root@tcp(192.168.92.128)/test")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	var wg sync.WaitGroup
	go scan_pr(db, q, &wg) // db is pointer already
	wg.Add(1)
	for i := 0; i < 1000; i++ {
		go check_pr(db, q, &wg)
		wg.Add(1)
	}
	wg.Wait() // time.Sleep(time.Second * 3600 * 10)
	fmt.Println("====task done==== All go routines finished executing")
}

func scan_pr(db *sql.DB, q chan *[100][2]string, wg *sync.WaitGroup) {
	start_id := ""
	t := 0
	for {
		fmt.Println("[scaning] from pr id: ", start_id)
		_sql := "SELECT pr.id, pr.property_id FROM public_record AS pr WHERE pr.id>'%s' ORDER BY pr.id ASC LIMIT 100;"
		// _sql := "SELECT a, b FROM A WHERE A.a>'%s' ORDER BY A.a ASC LIMIT 2;"
		rows, err := db.Query(strings.Replace(_sql, "%s", start_id, 1))
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		defer rows.Close()
		//--------------
		var _id sql.NullString
		var _p_id sql.NullString
		var id string
		var p_id string
		// -------------
		var l [100][2]string
		i := 0
		for rows.Next() {
			err = rows.Scan(&_id, &_p_id)
			if _id.Valid {
				id = _id.String
			} else {
				id = ""
			}
			if _p_id.Valid {
				p_id = _p_id.String
			} else {
				p_id = ""
			}
			if err != nil {
				fmt.Println(err)
				panic(err.Error())
			}
			// fmt.Println(id)
			l[i][0] = id
			l[i][1] = p_id
			i += 1
			start_id = id // ?
		}
		t += i
		fmt.Println("------------------------------------total= ", t)
		if l[0][0] == "" {
			fmt.Println("[done] scan pr done")
			for i := 0; i < 1000; i++ { //number = check_pr
				q <- &l
			}
			wg.Done()
			break
		}
		q <- &l
	}
}

func check_pr(db *sql.DB, q chan *[100][2]string, wg *sync.WaitGroup) {
	var c chan bool = make(chan bool)
	for {
		l := *(<-q)
		if l[0][0] == "" {
			fmt.Println("[exit] of 1 check_pr")
			wg.Done()
			break
		}
		fmt.Println("[checking] 100 pr")
		for _, i := range l {
			pr_id := i[0]
			pr_p_id := i[1]
			if pr_id == "" {
				break
			}
			// 1.
			if pr_p_id == "" {
				fmt.Println("[found] pr have no property: ", pr_id)
				continue
			}
			// 2.
			go find_if_p_exists(pr_p_id, db, c)
			pr_p_exists := <-c //need to block here in logic to wait another goroutine's result
			if !pr_p_exists {
				fmt.Println("[found] pr property not exists: ", pr_id)
				continue
			}
			// 3.
			go find_pr_multi_p(pr_id, db, c)
			<-c //need to block here to make sure this goroutine already finish
		}
	}
}

func find_if_p_exists(pr_p_id string, db *sql.DB, c chan bool) {
	sql_to_find_if_p_exist := "SELECT id FROM mls_public_record_association AS mp WHERE mp.id='{property_id}';"
	rows, err := db.Query(strings.Replace(sql_to_find_if_p_exist, "{property_id}", pr_p_id, 1))
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	var p_id string
	for rows.Next() { // if p_id not exists, won't come into this loop at all, so p_id == ""
		err = rows.Scan(&p_id)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
	}
	if p_id == "" {
		c <- false
	} else {
		c <- true
	}
}

func find_pr_multi_p(pr_id string, db *sql.DB, c chan bool) {
	sql_to_find_pr_multi_p := "SELECT id FROM mls_public_record_association AS mp WHERE mp.public_record_id='{pr_id}';"
	rows, err := db.Query(strings.Replace(sql_to_find_pr_multi_p, "{pr_id}", pr_id, 1))
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	var p_id string
	var p_id_list []string
	for rows.Next() { // definitely into this loop, just not sure how many times all.
		err = rows.Scan(&p_id)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		p_id_list = append(p_id_list, p_id)
	}
	if len(p_id_list) > 1 {
		fmt.Println("[found] pr has multi property: ", pr_id)
	}
	c <- false
}
