package handler

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)


type Hd struct {
	Db *sqlx.DB
}

func (hd *Hd) Test(c echo.Context) error {
	hd.Db = hd.Db;
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