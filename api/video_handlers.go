package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
)

// VideoHandler ハンドラ
type VideoHandler struct {
	VideoRepository model.VideoRepository
}

//GetVideos すべてのビデオ
func (h *VideoHandler) GetVideos(c echo.Context) error {
	videos := h.VideoRepository.FindAll()
	json, err := json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}

// SearchVideos 条件を指定して動画を探す
func (h *VideoHandler) SearchVideos(c echo.Context) error {
	channelIDparam := c.QueryParam("channel_id")
	channelIDs := strings.Split(channelIDparam, ",")
	if len(channelIDs) == 0 {
		channelIDs = append(channelIDs, channelIDparam)
	}
	videos := h.VideoRepository.FindByChannelID(channelIDs)
	json, err := json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}
