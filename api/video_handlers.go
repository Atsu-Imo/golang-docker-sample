package api

import (
	"os"
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"

	//postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//GetVideos すべてのビデオ
func GetVideos(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	db, err := gorm.Open("postgres", os.Getenv("DB"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	
	defer db.Close()
	var videos []model.Video
	db.Find(&videos)
	json, err :=json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}
