package server

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

)