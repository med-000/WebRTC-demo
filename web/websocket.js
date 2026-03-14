const ws = new WebSocket("ws://localhost:8080/ws");
//websocketサーバーがこのURLだよっていってるだけであって作っているわけじゃない。作っているのは接続のみ

ws.onopen = () => {
  console.log("ws connected");

  const payload = {
    message: "hello",
  };

  ws.send(JSON.stringify(payload));
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log(data);
};

ws.onclose = () => {
  console.log("closed");
};
