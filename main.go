package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
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
	e.GET("/comment/:meetingId", GetCommentFromId)
	e.GET("/reaction/:meetingId", GetReactionFromId)

	e.POST("/meeting", PostMeeting)
	e.GET("/meeting", GetMeeting)
	e.PATCH("/meeting/:meetingId", PatchMeetingFromId)
	e.GET("/meeting/:meetingId", GetMeetingFromId)
	e.DELETE("/meeting/:meetingId", DeleteMeetingFromId)

	e.POST("/token", PostToken)
	e.GET("/token", GetToken)
	e.GET("/token/:token", GetTokenFromToken)
	e.PATCH("/token/:token", PatchTokenFromToken)

	e.Logger.Fatal(e.Start(":8000"))
}

func GetCommentFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetReactionFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PostMeeting(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetMeeting(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PatchMeeting(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetMeetingFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PatchMeetingFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func DeleteMeetingFromId(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PostToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func GetTokenFromToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}

func PatchTokenFromToken(c echo.Context) error {
	return c.String(http.StatusNotImplemented, "未実装です")
}
