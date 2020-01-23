package api

import (
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
)

// VideoHandler ハンドラ
type VideoHandler struct {
	VideoRepository model.VideoRepository
}

//GetVideos すべてのビデオ
func (h *VideoHandler)GetVideos(c echo.Context) error {
	videos := h.VideoRepository.FindAll()
	json, err :=json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}

// GetVideosByChannelID チャンネルID指定
func (h *VideoHandler)GetVideosByChannelID(c echo.Context) error {
	videos := h.VideoRepository.FindByChannelID(c.QueryParam("channel_id"))
	json, err :=json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))

}