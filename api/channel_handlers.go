package api

import (
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
)

// ChannelHandler ハンドラ
type ChannelHandler struct {
	ChannelRepository model.ChannelRepository
}

//GetChannels すべてのチャンネル
func (h *ChannelHandler) GetChannels(c echo.Context) error {
	channels := h.ChannelRepository.FindAll()
	json, err :=json.Marshal(channels)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}

// GetChannelBy チャンネルIDを指定してチャンネル情報を取得する
func (h *ChannelHandler) GetChannelBy(c echo.Context) error {
	channelID := c.QueryParam("channel_id")
	if channelID == "" {
		return c.String(http.StatusOK, "")
	}
	channel := h.ChannelRepository.FindByChannelID(channelID)
	json, err :=json.Marshal(channel)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}