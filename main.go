package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
