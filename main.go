package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/traPtitech/Emoine_R/handler"
)

func main() {
	// TODO: 認証
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// デバッグ用 /debugを叩くと認証したものとみなす
	e.GET("/debug", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["userid"] = "s9"
		err := sess.Save(c.Request(), c.Response())
		if err != nil {
			return c.String(http.StatusInternalServerError, "セッションの保存に失敗しました")
		}

		return c.String(http.StatusOK, "あなたの名前を"+fmt.Sprint(sess.Values["userid"])+"として認証しました")
	})

	withLogin := e.Group("")
	withLogin.Use(handler.CheckLogin)

	withLogin.GET("/meeting", handler.GetMeeting)
	withLogin.GET("/meeting/:meetingId", handler.GetMeetingFromID)
	withLogin.GET("/meeting/:meetingId/comments", handler.GetCommentFromMeetingID)
	withLogin.GET("/meeting/:meetingId/reactions", handler.GetReactionFromMeetingID)

	withAdmin := withLogin.Group("")
	withAdmin.Use(handler.CheckIsAdmin)

	withAdmin.POST("/meeting", handler.PostMeeting)
	withAdmin.PATCH("/meeting/:meetingId", handler.PatchMeetingFromID)
	withAdmin.DELETE("/meeting/:meetingId", handler.DeleteMeetingFromID)
	withAdmin.POST("/token", handler.PostToken)
	withAdmin.GET("/token", handler.GetToken)
	withAdmin.GET("/token/:token", handler.GetTokenFromToken)
	withAdmin.PATCH("/token/:token", handler.PatchTokenFromToken)

	e.Logger.Fatal(e.Start(":8090"))
}
