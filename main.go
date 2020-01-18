package main

import (
	"github.com/joho/godotenv"
	"os"
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
	chanRepo := model.NewChannelRepository(os.Getenv("DB"))
	channelHandler := api.NewChannelHandler(chanRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", api.Hello)
	e.GET("/channels", channelHandler.GetChannels)
	e.GET("/channel", channelHandler.GetChannelBy)
	e.GET("/videos", api.GetVideos)
	e.Logger.Fatal(e.Start(":1323"))
}
