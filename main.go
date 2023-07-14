package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/traPtitech/Emoine_R/handler"
)

func main() {
	// TODO: 認証
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// デバッグ用 /debugを叩くと認証したものとみなす
	e.GET("/debug", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["userid"] = "aaa"
		err := sess.Save(c.Request(), c.Response())
		if err != nil {
			return c.String(http.StatusInternalServerError, "セッションの保存に失敗しました")
		}

		return c.String(http.StatusOK, "あなたの名前を"+fmt.Sprint(sess.Values["userid"])+"として認証しました")
	})

	e.POST("/oauth/generate/code", handler.OAuthGenerateCodeHandler)
	e.POST("/oauth/callback", handler.OAuthCallbackHandler)

	withLogin := e.Group("")
	withLogin.Use(handler.CheckLogin)

	withLogin.GET("/meeting", handler.GetMeetings)
	withLogin.GET("/meeting/:meetingId", handler.GetMeeting)
	withLogin.GET("/meeting/:meetingId/comments", handler.GetMeetingComments)
	withLogin.GET("/meeting/:meetingId/reactions", handler.GetMeetingReactions)

	withAdmin := withLogin.Group("")
	withAdmin.Use(handler.CheckIsAdmin)

	withAdmin.POST("/meeting", handler.CreateMeeting)
	withAdmin.PATCH("/meeting/:meetingId", handler.UpdateMeeting)
	withAdmin.DELETE("/meeting/:meetingId", handler.DeleteMeeting)
	withAdmin.GET("/meeting/:meetingId/tokens", handler.GetMeetingTokens)
	withAdmin.POST("/token", handler.CreateToken)
	withAdmin.GET("/token", handler.GetTokens)
	withAdmin.GET("/token/:token", handler.GetToken)
	withAdmin.PATCH("/token/:token", handler.UpdateToken)

	e.Logger.Fatal(e.Start(":8090"))
}
