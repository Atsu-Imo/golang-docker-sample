package main

import (
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func getChannels(c echo.Context) error {
	db, err := gorm.Open("postgres", "host=postgres port=5432 user=postgres password=password dbname=test_db sslmode=disable")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	
	defer db.Close()
	var channels []model.Channel
	db.Find(&channels)
	json, err :=json.Marshal(channels)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}
