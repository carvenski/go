package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	fmt.Println("start------")
	for i := 0; i < 10; i++ {
		// query_mysql(i)  //if you don't add go, it will just execute by order, no goroutines no concurrency !
		go query_mysql(i) // 所以说同一个go函数,既可以写普通顺序阻塞代码,也可以写协程异步非阻塞代码...就看前面加不加go
	}
	time.Sleep(time.Second * 10)
	fmt.Println("end------")
}

func query_mysql(i int) {
	fmt.Println("in goroutine------", i)
	db, err := sql.Open("mysql", "movoto:movoto123!@tcp(db3.ng.movoto.net)/movoto")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	//--------------------------------------------------------------------------------------------------------------
	//when io block(query db), this goroutine will yield CPU to let another goroutine use. just like python routine.
	//--------------------------------------------------------------------------------------------------------------
	rows, err := db.Query("select id from mls_listing order by id DESC limit 1")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var id string
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("get id: ", id)
		fmt.Println("out goroutine------", i)
	}
}
