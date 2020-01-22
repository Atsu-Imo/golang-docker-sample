package api

import (
	"net/http"
	"encoding/json"

	"github.com/Atsu-Imo/golang-docker-sample/service"
	"github.com/labstack/echo/v4"
)

// ChannelHandler ハンドラ
type VideoHandler struct {
	service service.VideoRepository
}
// NewChannelHandler 本来はinterfaceを返却すべきなので余裕があったら修正
func NewVideoHandler(c service.VideoRepository) *VideoHandler {
	return &VideoHandler{service: c}
}

//GetVideos すべてのビデオ
func (h *VideoHandler) GetVideos(c echo.Context) error {
	videos := h.service.FindAll()
	json, err :=json.Marshal(videos)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, string(json))
}
