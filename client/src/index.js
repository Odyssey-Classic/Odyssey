const messages = require('../pb/game_message_pb.js')

let ws = new WebSocket('ws://172.18.87.126:3001/ws');
ws.binaryType = "arraybuffer";
ws.onopen = (event) => {
    var message = new messages.GameMessage();
    message.setType(1);
    var bytes = message.serializeBinary();
    console.log(message.getType());
    console.log(bytes);
    ws.send(bytes);
};

ws.addEventListener('error', (event) => {
    console.error('Error:', event);
});

ws.addEventListener('close', (event) => {
    console.log('Disconnected from server');
});

ws.addEventListener('message', (event) => {
    console.log('Message:', event.data);
})
