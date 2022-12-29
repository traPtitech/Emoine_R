package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var (
	options = sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &options;

		if sess.Values["userid"] == nil {
			return c.String(http.StatusForbidden, "ログインしてください")
		}

		return next(c)
	}
}

func CheckIsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &options;

		adminNames := strings.Split(os.Getenv("ADMIN_NAMES"), ",")
		myname := sess.Values["userid"]
		isAdmin := false
		for _, adminName := range adminNames {
			if myname == adminName {
				isAdmin = true
			}
		}
		
		if !isAdmin {
			return c.String(http.StatusForbidden, "権限がありません")
		}

		return next(c)
	}
}