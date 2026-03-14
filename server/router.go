package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/ws",WsHandler)

	http.ListenAndServe(":8080", nil)
}