package api

import(
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expect := "Hello World!"
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("hello")
	if assert.NoError(t, Hello(c)) {
		assert.Equal(t, expect, rec.Body.String())
	}
}