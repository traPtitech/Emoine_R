package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReactionFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
