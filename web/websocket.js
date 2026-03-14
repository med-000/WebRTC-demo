const ws = new WebSocket(`ws://localhost:8080/room/${id}`);
//websocketサーバーがこのURLだよっていってるだけであって作っているわけじゃない。作っているのは接続のみ

ws.onopen = () => {
  console.log("ws connected");

  const signalMessage = {
    type: "",
    sdp: "",
    candidate: "",
    sdpMid: "",
    sdpMLineIndex: 0,
  };

  ws.send(JSON.stringify(signalMessage));
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log(data);
};

ws.onclose = () => {
  console.log("ws closed");
};
