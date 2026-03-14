package server

import (
	"net/http"
)

func Start() {

	// static files (js css)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("web/static")),
		),
	)

	http.HandleFunc("/", Root)
	http.HandleFunc("/room/", RoomHandler)
	http.HandleFunc("/create", CreateRoom)
	http.HandleFunc("/join", JoinRoom)
	http.HandleFunc("/ws/", WsHandler)

	http.ListenAndServe(":8080", nil)
}