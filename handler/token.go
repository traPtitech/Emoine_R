package handler

import (
	"time"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/traPtitech/Emoine_R/model"
	"github.com/traPtitech/Emoine_R/model/dbschema"
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
	tokenDBSchema, err := dbschema.TokenByToken(c.Request().Context(), model.DB, c.Param("token"))

	if err != nil {
		return c.String(http.StatusNotFound, "tokenが見つかりませんでした: "+err.Error())
	}
	token := schema.Token{
			CreatedAt: 		tokenDBSchema.CreatedAt,
			CreatorId: 		tokenDBSchema.CreatorID,
			Description: 	mustValue[string](tokenDBSchema.Description),
			ExpireAt: 		mustValue[time.Time](tokenDBSchema.ExprieAt),
			MeetingId: 		tokenDBSchema.MeetingID,
			Token: 			tokenDBSchema.Token,
			Username: 		tokenDBSchema.UserID,
	}

	return c.JSON(http.StatusOK, token)
}

func PatchTokenFromToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
