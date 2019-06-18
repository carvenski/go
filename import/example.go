package main

import (
	// 这里使用的是相对路径,这样写才可以认识当前路径的包...
	// 要么就把根目录加入GOPATH
    api "./api"    
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()
    e.GET("/", api.HelloWorld)
    e.Logger.Fatal(e.Start(":1323"))
}

