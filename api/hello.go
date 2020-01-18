package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

//Hello サンプルとしておいておく
func Hello(c echo.Context) error{
	return c.String(http.StatusOK, "Hello World!")
}