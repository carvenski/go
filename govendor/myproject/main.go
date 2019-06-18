package main

import (
    "log"
    "net/http"
    "myproject/lib"

    "github.com/labstack/echo"
)

func main(){
    log.Println("hello")
    lib.Hello()

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
       return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}


