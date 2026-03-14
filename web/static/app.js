// ----------------------------
// roomId取得
// ----------------------------
const roomId = window.location.pathname.split("/").pop();

// ----------------------------
// WebSocket
// ----------------------------
const ws = new WebSocket(`ws://${location.host}/ws/${roomId}`);

// ----------------------------
// WebRTC設定
// ----------------------------
const config = {
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
};

let pc = null;
let dc = null;

// ----------------------------
// signaling送信
// ----------------------------
function sendSignal(data) {
  ws.send(JSON.stringify(data));
}

// ----------------------------
// DataChannel
// ----------------------------
function setupDataChannel(channel) {
  channel.onopen = () => {
    console.log("DataChannel open");
  };

  channel.onmessage = (event) => {
    const history = document.getElementById("history");

    history.value = "other> " + event.data + "\n" + history.value;
  };
}

// ----------------------------
// PeerConnection作成
// ----------------------------
function createPeerConnection() {
  pc = new RTCPeerConnection(config);

  pc.onicecandidate = (event) => {
    if (!event.candidate) return;

    sendSignal({
      type: "candidate",
      candidate: event.candidate,
    });
  };

  pc.ondatachannel = (event) => {
    dc = event.channel;

    setupDataChannel(dc);
  };
}

// ----------------------------
// offer作成
// ----------------------------
async function startOffer() {
  createPeerConnection();

  dc = pc.createDataChannel("chat");

  setupDataChannel(dc);

  const offer = await pc.createOffer();

  await pc.setLocalDescription(offer);

  sendSignal({
    type: "offer",
    sdp: offer.sdp,
  });
}

// ----------------------------
// offer受信
// ----------------------------
async function handleOffer(sdp) {
  createPeerConnection();

  await pc.setRemoteDescription({
    type: "offer",
    sdp: sdp,
  });

  const answer = await pc.createAnswer();

  await pc.setLocalDescription(answer);

  sendSignal({
    type: "answer",
    sdp: answer.sdp,
  });
}

// ----------------------------
// answer受信
// ----------------------------
async function handleAnswer(sdp) {
  await pc.setRemoteDescription({
    type: "answer",
    sdp: sdp,
  });
}

// ----------------------------
// candidate受信
// ----------------------------
async function handleCandidate(msg) {
  if (!pc) return;

  await pc.addIceCandidate(msg.candidate);
}

// ----------------------------
// chat送信
// ----------------------------
function sendMessage() {
  const input = document.getElementById("message");
  const history = document.getElementById("history");

  const msg = input.value;

  input.value = "";

  dc.send(msg);

  history.value = "me> " + msg + "\n" + history.value;
}

window.sendMessage = sendMessage;

// ----------------------------
// WebSocket
// ----------------------------
ws.onopen = () => {
  console.log("ws connected");

  // デモでは最初の接続者がoffer
  startOffer();
};

ws.onmessage = async (event) => {
  const msg = JSON.parse(event.data);

  switch (msg.type) {
    case "offer":
      await handleOffer(msg.sdp);
      break;

    case "answer":
      await handleAnswer(msg.sdp);
      break;

    case "candidate":
      await handleCandidate(msg);
      break;
  }
};
