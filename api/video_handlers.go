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
	channelIDs := strings.Split(c.QueryParam("channel_id"), ",")
	videos := h.VideoRepository.FindByChannelID(channelIDs)
	json, err := json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}
