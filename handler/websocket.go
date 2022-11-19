package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	// TODO: このCORS回避部分は全オリジンを通しているが、本番環境の場合に変更するようにしたい
	CheckOrigin: func(r *http.Request) bool {
		return true
  },
}

func WebSocketHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if (err != nil) {
		log.Printf("Error when updating the connection to ws: %s", err)
		return err;
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if (err != nil) {
			log.Printf("Error when reading message in websocket: %s", err)
			return err;
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error when sending message in websocket: %s", err)
			return err;
		}

		// TODO: メッセージの構文解析（JSONが好ましい？）をして文章あるいはスタンプを取得し、以下の関数に投げる
	  // go WSMessageHandler(meetingId, message);
	  // go WSStampHandler(meetingId, stampId);
	}
}

func WSMessageHandler() {
  // TODO
	// メッセージをDBにPushする
	// 全員に同一メッセージを送る
}
func WSStampHandler() {
  // TODO
	// スタンプをDBにPushする
	// 全員に同一スタンプを送る
}