package main

import (
	"fmt"
	"log"
	"os"

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

	e.GET("/comment/:meetingId", handler.GetCommentFromId)
	e.GET("/reaction/:meetingId", handler.GetReactionFromId)

	e.POST("/meeting", handler.PostMeeting)
	e.GET("/meeting", handler.GetMeeting)
	e.PATCH("/meeting/:meetingId", handler.PatchMeetingFromId)
	e.GET("/meeting/:meetingId", handler.GetMeetingFromId)
	e.DELETE("/meeting/:meetingId", handler.DeleteMeetingFromId)

	e.POST("/token", handler.PostToken)
	e.GET("/token", handler.GetToken)
	e.GET("/token/:token", handler.GetTokenFromToken)
	e.PATCH("/token/:token", handler.PatchTokenFromToken)

	// Notice: WebSocketはポート443で運用するのがいいかも？
	e.GET("/ws", handler.WebSocketHandler)
	
	e.Logger.Fatal(e.Start(":8090"))
}







