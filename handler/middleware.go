package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var (
	envError error = godotenv.Load()
	ClientID string = os.Getenv("CLIENT_ID")
	SessionKey string = "session"

	SessionOptionsDefault sessions.Options = sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 1000,
		HttpOnly: true,
	}
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(SessionKey, c)
		sess.Options = &SessionOptionsDefault;

		if sess.Values["userid"] == nil {
			return c.String(http.StatusForbidden, "ログインしてください")
		}

		return next(c)
	}
}

func CheckIsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(SessionKey, c)
		sess.Options = &SessionOptionsDefault;

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
