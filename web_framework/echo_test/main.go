package main

import (
	"myapp/app"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	app.InitRouter(e)
	e.Logger.Fatal(e.Start(":30000"))
}
