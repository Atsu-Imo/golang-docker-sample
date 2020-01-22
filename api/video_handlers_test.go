package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/Atsu-Imo/golang-docker-sample/service"
	"github.com/stretchr/testify/assert"
)

func TestGetVideos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("videos")
	c.QueryParams().Add("channel_id", "test")
	r := service.VideoRepository{}
	h := NewVideoHandler(r)
	if assert.NoError(t, h.GetVideos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, nil, rec.Body.String())
	}
}
