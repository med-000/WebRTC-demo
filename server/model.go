package server

import "github.com/gorilla/websocket"

type (
SignalMessage struct {
	Type string `json:"type"`

	SDP string `json:"sdp,omitempty"`

	Candidate string `json:"candidate,omitempty"`
	SdpMid string `json:"sdpMid,omitempty"`
	SdpMLineIndex int `json:"sdpMLineIndex,omitempty"`
}

Client struct {
	Conn *websocket.Conn
	Room string
}

HTMLtemplate struct {
	Title string
}
)