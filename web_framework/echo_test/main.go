package main

import (
	"myapp/app"

	//使用echo v4版本
	//参考https://echo.labstack.com/guide/
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	// init url router
	app.InitRouter(e)
	// server start
	e.Logger.Fatal(e.Start(":30000"))
}
