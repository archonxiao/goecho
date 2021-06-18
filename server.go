package main

import (
	"github.com/labstack/echo/v4"
	"goecho/dao"
	"goecho/router"
)

func main() {
	dao.Connect()
	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":1234"))
}
