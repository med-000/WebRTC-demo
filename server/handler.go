package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var rooms = map[string][]*Client{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	fmt.Println("client connected")
	client := &Client{
		Conn: conn,
	}

	fmt.Println("client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var m Message
		json.Unmarshal(msg, &m)

		switch m.Type {

		case "join":

			client.Room = m.Room

			rooms[m.Room] = append(rooms[m.Room], client)

			fmt.Println("client joined room:", m.Room)

		case "offer", "answer", "ice":

			fmt.Println("relay message:", m.Type)

			for _, c := range rooms[m.Room] {

				if c.Conn != conn {

					c.Conn.WriteMessage(websocket.TextMessage, msg)

				}

			}

		}
	}
}