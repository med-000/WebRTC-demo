package server

import "github.com/gorilla/websocket"

type (

Signal struct {
	Type string `json:"type"`
	Room string `json:"room"`
	Status string `json:"status"`
	Log string `json:"log"`

	SDP string `json:"sdp,omitempty"`
	Candidate string `json:"candidate,omitempty"`
	SdpMid string `json:"sdpMid,omitempty"`
	SdpMLineIndex int `json:"sdpMLineIndex,omitempty"`
}
Message struct {
	Type string `json:"type"`
	Room string `json:"room"`
	Data string `json:"data,omitempty"`
}

Client struct {
	Conn *websocket.Conn
	Room string
}
)