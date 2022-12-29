package handler

import (
	"context"
	"net/http"

	session "github.com/go-session/session/v3"
	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		store, err := session.Start(context.Background(), c.Response(), c.Request());
		if err != nil {
			return c.String(http.StatusInternalServerError, "セッション情報が読み込めません")
		}
		_, ok := store.Get("userid")

		if !ok {
			return c.String(http.StatusForbidden, "ログインしてください")
		}
		return next(c);
	}
}