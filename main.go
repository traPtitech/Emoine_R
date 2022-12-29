package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	session "github.com/go-session/session/v3"
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

	// デバッグ用 /debugを叩くと認証したものとみなす
	e.GET("/debug", func(c echo.Context) error{
		store, err := session.Start(c.Request().Context(), c.Response(), c.Request());
		if err != nil {
			return c.String(http.StatusInternalServerError, "セッション情報が読み込めません")
		}
		store.Set("userid", "testid")
		err = store.Save()
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprint(err))
		}
		return c.String(http.StatusOK, "あなたの名前をtestとして認証しました")
	})

	// なろう講習会
	// session + basic認証（ユーザー名＋パスワードを入力）
	// サーバー上にusername

	// session/ + OAuth認証（+ token認証）

	withLogin := e.Group("")
	withLogin.Use(handler.SessionMiddleware)
	
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