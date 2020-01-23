package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Atsu-Imo/golang-docker-sample/api"
	"github.com/Atsu-Imo/golang-docker-sample/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open("postgres", os.Getenv("DB"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	chanRepo := model.NewChannelRepository(db)
	channelHandler := &api.ChannelHandler{ChannelRepository: chanRepo}

	videoRepo := model.NewVideoRepository(db)
	videoHandler := &api.VideoHandler{VideoRepository: videoRepo}
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", api.Hello)
	e.GET("/channels", channelHandler.GetChannels)
	e.GET("/channel", channelHandler.GetChannelBy)
	e.GET("/videos", videoHandler.GetVideos)
	e.Logger.Fatal(e.Start(":1323"))
}
