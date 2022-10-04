var WebSocket = require ('ws').WebSocket;

const ws = new WebSocket('ws://localhost:8888/ws',
    {
        headers: {"AAA": "=== AAA header ==="}, 
        host: "=== AAA.com ===" 
    }
);

console.log(ws)

ws.on('open', function open() {
  console.log("=== ws open")
});
ws.on('message', function incoming(message) {
  console.log('=== ws get msg: %s', message);
});
