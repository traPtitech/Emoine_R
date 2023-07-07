package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/traPtitech/Emoine_R/handler/schema"
)

func GetToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PostToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

// stringなキーのtokenから構造体のtokenを得る。
func GetTokenFromToken(c echo.Context) error {
	token_string := c.Param("token")
	var token_struct schema.Token
	if err := model.Db.Get(&token_struct, "SELECT * FROM token WHERE token='"+token_string+"'"); errors.Is(err, sql.ErrNoRows) {
		return c.String(http.StatusNotFound, "tokenが見つかりませんでした: "+err.Error())
	}
	return c.JSON(http.StatusOK, token_struct)
}

func PatchTokenFromToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
