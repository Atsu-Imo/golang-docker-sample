package main

import (
	"github.com/labstack/echo/v4"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main(){
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))

	db, err := gorm.Open("postgres", "host=myhost port=5432 user=root password=postgres")
	if err != nil {
		panic(err)
	}
	fmt.Println(db.Debug())
	defer db.Close()
}
