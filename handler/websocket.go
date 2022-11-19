// これを読んでいる自分へ
// 新しいコードを書く負担よりは、siguredoを使うことにした
// スタンプとコメントをDBに保存する際、meetingIDが必要で、そのためにはsession管理が必要

package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shiguredo/websocket"
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
		myString := string(p[:]);
		log.Printf("messageType: " + strconv.Itoa(messageType) + " p: " + myString);

		// TODO: メッセージの構文解析をして文章あるいはスタンプを取得し、以下の関数に投げる
	  // go WSMessageHandler(conn, meetingId, message);
	  // go WSStampHandler(conn, meetingId, stampId);
		
		// 構文解析：先頭の文字がMかSで調べるくらいでよさそう
	}
}

func WSMessageHandler(meetingId string, message string) {
	log.Printf("WSMessageHandler(" + meetingId + "," + message + ")");
  // TODO
	// メッセージをDBにPushする
	// 全員に同一メッセージを送る
}
func WSStampHandler(meetingId string, stampId string) {
	log.Printf("WSMessageHandler(" + meetingId + "," + stampId + ")");
  // TODO
	// スタンプをDBにPushする
	// 全員に同一スタンプを送る
}