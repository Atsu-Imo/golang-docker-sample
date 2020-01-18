package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Atsu-Imo/golang-docker-sample/api"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/hello", api.Hello)
	e.GET("/channels", api.GetChannels)
	e.GET("/channel", api.GetChannelByChannelID)
	e.GET("/videos", api.GetVideos)
	e.Logger.Fatal(e.Start(":1323"))
}
