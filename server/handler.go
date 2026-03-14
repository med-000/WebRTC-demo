package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/websocket"
)
var joincreatetemp = template.Must(template.ParseFiles("web/templates/joincreate.html"))
var roomtemp = template.Must(template.ParseFiles("web/templates/room.html"))

var rooms = map[string][]*Client{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Root(w http.ResponseWriter, r *http.Request) {

	data := HTMLtemplate{
		Title:"Home",
	}

	joincreatetemp.Execute(w,data)
}

func RoomHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Path[len("/room/"):]
	data := HTMLtemplate{
		Title:"Room" + room,
	}
	roomtemp.Execute(w,data)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(1000)
	
	url := fmt.Sprintf("/room/%d", n)

	http.Redirect(w, r, url, http.StatusFound)
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {

	roomIdStr := r.FormValue("roomId")

	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		http.Error(w, "invalid room id", http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("/room/%d", roomId)

	http.Redirect(w, r, url, http.StatusFound)
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Path[len("/ws/"):]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	fmt.Println("client connected")

	client := &Client{
		Conn: conn,
		Room: room,
	}

	rooms[room] = append(rooms[room], client)

	fmt.Println("joined room:",room)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		for _, c := range rooms[room] {
			if c.Conn != conn {

				c.Conn.WriteMessage(websocket.TextMessage, msg)

			}
		}		
	}
}