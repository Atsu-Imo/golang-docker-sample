package api

import (
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
)

type handler struct {
	channelRepository model.ChannelRepository
}
func NewChannelHandler(c model.ChannelRepository) *handler {
	return &handler{channelRepository: c}
}

//GetChannels すべてのチャンネル
func (h handler) GetChannels(c echo.Context) error {
	channels := h.channelRepository.FindAll()
	json, err :=json.Marshal(channels)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}

// GetChannelBy チャンネルIDを指定してチャンネル情報を取得する
func (h handler) GetChannelBy(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusOK, "")
	}
	channel := h.channelRepository.FindByChannelID(channelID)
	json, err :=json.Marshal(channel)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}