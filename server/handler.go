package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
}
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	fmt.Println("client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println(string(msg))

		conn.WriteMessage(websocket.TextMessage, msg)
	}
}