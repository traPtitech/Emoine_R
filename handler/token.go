package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
)

func GetToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PostToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

// stringなキーのtokenから構造体のtokenを得る。
func GetTokenFromToken(c echo.Context) error {
	tokenPrimaryKey := c.Param("token")
	tokenRecord, err := dbschema.TokenByToken(c.Request().Context(), model.DB, tokenPrimaryKey)

	if err != nil {
		return c.String(http.StatusNotFound, "tokenが見つかりませんでした: "+err.Error())
	}
	return c.JSON(http.StatusOK, tokenRecord)
}

func PatchTokenFromToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
