package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type stubVideoRepository struct {
}

func (r *stubVideoRepository) FindAll() []model.Video {
	video1 := model.Video{ChannelID: "channel1"}
	video2 := model.Video{ChannelID: "channel2"}
	return []model.Video{video1, video2}
}

func (r *stubVideoRepository) FindByChannelID(channelID []string, page int, limit int) []model.Video {
	video1 := model.Video{ChannelID: channelID[0]}
	video2 := model.Video{ChannelID: channelID[0]}
	return []model.Video{video1, video2}
}

func TestGetVideos(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("videos")
	r := &stubVideoRepository{}
	h := VideoHandler{VideoRepository: r}
	if assert.NoError(t, h.GetVideos(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var videos []model.Video
		json.Unmarshal(rec.Body.Bytes(), &videos)
		video1 := videos[0]
		video2 := videos[1]
		assert.Equal(t, "channel1", video1.ChannelID)
		assert.Equal(t, "channel2", video2.ChannelID)
	}
}

func TestSearchVideo(t *testing.T) {
	assert.Fail(t, "TODO")
}
