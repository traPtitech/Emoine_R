package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/traPtitech/Emoine_R/handler"
)

var (
	Db *sqlx.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot collect .env: %s", err)
	}
	_db, err := sqlx.Connect(
		"mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/%s", os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	Db = _db

	// TODO: 認証
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// デバッグ用 /debugを叩くと認証したものとみなす
	e.GET("/debug", func(c echo.Context) error{
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
	
	withLogin.GET("/comment/:meetingId", handler.GetCommentFromId)
	withLogin.GET("/reaction/:meetingId", handler.GetReactionFromId)
	withLogin.GET("/meeting", handler.GetMeeting)
	withLogin.GET("/meeting/:meetingId", handler.GetMeetingFromId)

	withAdmin := withLogin.Group("")
	withAdmin.Use(handler.CheckIsAdmin)

	withAdmin.POST("/meeting", handler.PostMeeting)
	withAdmin.PATCH("/meeting/:meetingId", handler.PatchMeetingFromId)
	withAdmin.DELETE("/meeting/:meetingId", handler.DeleteMeetingFromId)
	withAdmin.POST("/token", handler.PostToken)
	withAdmin.GET("/token", handler.GetToken)
	withAdmin.GET("/token/:token", handler.GetTokenFromToken)
	withAdmin.PATCH("/token/:token", handler.PatchTokenFromToken)


	// Notice: WebSocketはポート443で運用するのがいいかも？
	e.GET("/ws", handler.WebSocketHandler)
	
	e.Logger.Fatal(e.Start(":8090"))
}
