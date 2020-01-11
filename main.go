package main

import (
	"github.com/labstack/echo/v4"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/Atsu-Imo/golang-docker-sample/model"
)

func main(){
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=test_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	var user model.User
	db.Debug().First(&user)
	fmt.Print(&user)
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))

}
