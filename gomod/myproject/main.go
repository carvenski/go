package main

import (
    "log"
    "net/http"
                                   //关于go mod支持包的不同版本的写法: /vN后缀
    //"github.com/labstack/echo"   //echo包的这个路径只支持到v3.3.10
    "github.com/labstack/echo/v4"  //加了v4后缀的这个路径才能支持v4.0.0以上的版本 (算个go mod的坑吧)
)

func main(){
    log.Println("hello")

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
       return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}


