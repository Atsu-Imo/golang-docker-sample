package api

import (
	"os"
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

//GetChannels すべてのチャンネル
func GetChannels(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	db, err := model.Connect(os.Getenv("DB"))
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

// GetChannelByChannelID チャンネルIDを指定してチャンネル情報を取得する
func GetChannelByChannelID(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusOK, "")
	}
	err := godotenv.Load()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	db, err := model.Connect(os.Getenv("DB"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	
	defer db.Close()
	var channels []model.Channel
	db.Where("channel_id = ?", channelID).Find(&channels)
	json, err :=json.Marshal(channels)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}