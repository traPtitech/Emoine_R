package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
	emoine_rv1 "github.com/traPtitech/Emoine_R/pkg/pbgen/emoine_r/v1"
)

func (h *AdminAPIHandler) GetTokens(ctx context.Context, req *connect.Request[emoine_rv1.GetTokensRequest]) (*connect.Response[emoine_rv1.GetTokensResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("未実装です"))
}

func GetTokens(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func CreateToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

// stringなキーのtokenから構造体のtokenを得る。
func GetToken(c echo.Context) error {
	token_string := c.Param("token")
	token_struct, err := dbschema.TokenByToken(c.Request().Context(), model.DB, token_string)
	if err != nil {
		return c.String(http.StatusNotFound, "tokenが見つかりませんでした: "+err.Error())
	}
	return c.JSON(http.StatusOK, token_struct)
}

func UpdateToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
