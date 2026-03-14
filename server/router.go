package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/room/",WsHandler)
	http.HandleFunc("/create",CreateRoom)
	http.HandleFunc("/join",JoinRoom)

	http.ListenAndServe(":8080", nil)
}