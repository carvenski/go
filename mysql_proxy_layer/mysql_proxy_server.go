package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "log"
	"net/http"
)

func main() {
	// mysql config
	host := flag.String("host", "localhost", "mysql host")
	port := flag.String("port", ":3306", "mysql port")
	user := flag.String("user", "root", "mysql user")
	passwd := flag.String("passwd", "", "mysql passwd")
	db := flag.String("db", "test", "mysql db")
	flag.Parse()

	params := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=60s", user, passwd, host, port, db)
	mysql_conn_pool, err := sql.Open("mysql", params) // mysql_conn_pool is a connection pool, pointer already
	// mysql_conn_pool.SetMaxIdleConns(10)               // set connection num in pool
	// mysql_conn_pool.SetMaxOpenConns(10)
	defer mysql_conn_pool.Close()
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/test", test)
	e.GET("/execute", execute)
	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:10000"))
}

// Handler
func test(c echo.Context) error {
	return c.String(http.StatusOK, "test ok !")
}

func execute(c echo.Context) error {
	return c.String(http.StatusOK, "test ok !")
}
