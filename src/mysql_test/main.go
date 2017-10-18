package main
import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	fmt.Println("start------")
	for i := 0; i < 10; i++{
	    go query_mysql(i)
	}
	time.Sleep(time.Second*10)
	fmt.Println("end------")
}

func query_mysql(i int){
	fmt.Println("in goroutine------", i)
	db, err := sql.Open("mysql", "movoto:movoto123!@tcp(db3.ng.movoto.net)/movoto?charset=utf8")
	if err != nil{fmt.Println(err)}
	defer db.Close()
	err = db.Ping()
	if err != nil{fmt.Println(err)}
        //--------------------------------------------------------------------------------------------------------------
        //when io block(query db), this goroutine will yield CPU to let another goroutine use. just like python routine.
        //-------------------------------------------------------------------------------------------------------------
	rows, err := db.Query("select id from mls_listing order by id DESC limit 1")
	if err != nil{fmt.Println(err)}
	defer rows.Close()
	var id string
	for rows.Next(){
	err = rows.Scan(&id)
	if err != nil{fmt.Println(err)}
	fmt.Println("get id: ", id)
	fmt.Println("out goroutine------", i)
	}
}


