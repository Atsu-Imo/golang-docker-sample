package api

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/stretchr/testify/assert"
)
type stubChannelRepository struct {
}
func newStub() stubChannelRepository{
	return stubChannelRepository{}
}
func (r stubChannelRepository) FindAll()[]model.Channel {
	return nil
}
func (r stubChannelRepository) FindByChannelID(channelID string) model.Channel {
	channel := model.Channel{ChannelID: channelID}
	return channel
}
func TestGetChannelBy(t *testing.T) {
	e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("channel")
	c.SetParamNames("channel_id")
	c.SetParamValues("test")
	c.QueryParams().Add("channel_id", "test")
	r := newStub()
	h := NewChannelHandler(r)
	if assert.NoError(t, h.GetChannelBy(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "testy", rec.Body.String())
	}
}