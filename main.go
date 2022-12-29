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
		sess.Values["userid"] = "SlimySlime"
		err := sess.Save(c.Request(), c.Response())
		if err {
			return c.String(http.StatusInternalServerError, "セッションの保存に失敗しました")
		}

		return c.String(http.StatusOK, "あなたの名前をSlimySlimeとして認証しました")
	})

	// なろう講習会
	// session + basic認証（ユーザー名＋パスワードを入力）
	// サーバー上にusername

	// session/ + OAuth認証（+ token認証）

	withLogin := e.Group("")
	withLogin.Use(checkLogin)
	
	withLogin.GET("/comment/:meetingId", handler.GetCommentFromId)
	withLogin.GET("/reaction/:meetingId", handler.GetReactionFromId)

	withLogin.POST("/meeting", handler.PostMeeting)
	withLogin.GET("/meeting", handler.GetMeeting)
	withLogin.PATCH("/meeting/:meetingId", handler.PatchMeetingFromId)
	withLogin.GET("/meeting/:meetingId", handler.GetMeetingFromId)
	withLogin.DELETE("/meeting/:meetingId", handler.DeleteMeetingFromId)

	withLogin.POST("/token", handler.PostToken)
	withLogin.GET("/token", handler.GetToken)
	withLogin.GET("/token/:token", handler.GetTokenFromToken)
	withLogin.PATCH("/token/:token", handler.PatchTokenFromToken)

	e.Logger.Fatal(e.Start(":8090"))
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		if sess.Values["userid"] == nil {
			return c.String(http.StatusForbidden, "ログインしてください")
		}
		log.Println(""+fmt.Sprint(sess.Values["userid"])+"が入りました")

		return next(c)
	}
}