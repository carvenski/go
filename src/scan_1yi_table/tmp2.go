package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"sync"
)

func main() {
	fmt.Println("[start] scan property table")
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
	go scan_property(db, q, &wg) // db is pointer already
	wg.Add(1)
	for i := 0; i < 1000; i++ {
		go check_property(db, q, &wg)
		wg.Add(1)
	}
	wg.Wait() // time.Sleep(time.Second * 3600 * 10)
	fmt.Println("====task done====")
}

func scan_property(db *sql.DB, q chan *[100][2]string, wg *sync.WaitGroup) {
	start_id := ""
	t := 0
	for {
		fmt.Println("[scaning] from property id: ", start_id)
		_sql := "SELECT mp.id, mp.public_record_id FROM mls_public_record_association AS mp WHERE mp.id>'%s' ORDER BY mp.id ASC LIMIT 100;"
		rows, err := db.Query(strings.Replace(_sql, "%s", start_id, 1))
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		defer rows.Close()
		//--------------
		var _id sql.NullString
		var _p_pr_id sql.NullString
		var id string
		var p_pr_id string
		// -------------
		var l [100][2]string
		i := 0
		for rows.Next() {
			err = rows.Scan(&_id, &_p_pr_id)
			if _id.Valid {
				id = _id.String
			} else {
				id = ""
			}
			if _p_pr_id.Valid {
				p_pr_id = _p_pr_id.String
			} else {
				p_pr_id = ""
			}
			if err != nil {
				fmt.Println(err)
				panic(err.Error())
			}
			l[i][0] = id
			l[i][1] = p_pr_id
			i += 1
			start_id = id // ?
		}
		t += i
		fmt.Println("------------------------------------total= ", t)
		if l[0][0] == "" {
			fmt.Println("[done] scan property done")
			for i := 0; i < 1000; i++ { //number = check_property
				q <- &l
			}
			wg.Done()
			break
		}
		q <- &l
	}
}

func check_property(db *sql.DB, q chan *[100][2]string, wg *sync.WaitGroup) {
	var c chan bool = make(chan bool)
	for {
		l := *(<-q)
		if l[0][0] == "" {
			fmt.Println("[exit] of 1 check_property")
			wg.Done()
			break
		}
		fmt.Println("[checking] 100 property")
		for _, i := range l {
			p_id := i[0]
			p_pr_id := i[1]
			if p_pr_id == "" { // skip p_pr_id not exists
				continue
			}
			if p_id == "" {
				break
			}
			go find_p_pr_exists_or_point_other(p_pr_id, p_id, db, c)
			<-c //need to block here to make sure this goroutine already finish
		}
	}
}

func find_p_pr_exists_or_point_other(pr_id string, pA_id string, db *sql.DB, c chan bool) {
	_sql := "SELECT id, property_id FROM public_record AS pr WHERE pr.id='{pr_id}';"
	rows, err := db.Query(strings.Replace(_sql, "{pr_id}", pr_id, 1))
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	var id string
	var _pB_id sql.NullString // _pB_id may be empty
	var pB_id string
	var c2 chan bool = make(chan bool)
	for rows.Next() { // if id not exists, won't into this loop
		err = rows.Scan(&id, &_pB_id)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		if _pB_id.Valid {
			pB_id = _pB_id.String
		} else {
			pB_id = ""
		}
		if id == "" {
			// 4.
			fmt.Println("[find] property's pr_id not exist, p_id: ", pA_id)
			go find_pr_by_p_address(pA_id, db, c2)
			<-c2 //need to block here to make sure this goroutine already finish
		} else {
			// 5.
			if pB_id != "" && pB_id != pA_id {
				fmt.Println("[find] p-A has pr has p-B ", pA_id, id, pB_id)
			}
		}
	}
	c <- false
}

func find_pr_by_p_address(p_id string, db *sql.DB, c chan bool) {
	_sql := `SELECT pr.id FROM mls_public_record_association AS mp 
             INNER JOIN address AS ad1 ON mp.address_id=ad1.id 
             INNER JOIN address AS ad2 ON ad1.address=ad2.address
             INNER JOIN public_record AS pr ON pr.address_id=ad2.id
             WHERE mp.id='{property_id}' 
             ORDER BY pr.update_time DESC 
             LIMIT 1 ;
             `
	rows, err := db.Query(strings.Replace(_sql, "{property_id}", p_id, 1))
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			panic(err.Error())
		}
		if id != "" {
			fmt.Println("[found] pr by p address: ", id)
		}
	}
	c <- false
}
