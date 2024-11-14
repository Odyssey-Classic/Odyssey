// This was prototype code that did basic WebSocket and ProtoBuf messaging stuff.
// Here for reference when we come back to do this more betterer.

// import { GameMessage } from '../pb/game_message'

// let ws = new WebSocket('ws://172.18.87.126:3001/ws')
// ws.binaryType = "arraybuffer"
// ws.onopen = (event) => {
//     var message = new GameMessage()
//     message.type = 1
//     var bytes = message.serializeBinary()
//     console.log(message.type)
//     console.log(bytes)
//     ws.send(bytes)
// }

// ws.addEventListener('error', (event) => {
//     console.error('Error:', event)
// })

// ws.addEventListener('close', (event) => {
//     console.log('Disconnected from server')
// })

// ws.addEventListener('message', (event) => {
//     console.log('Message:', event.data)
//     var msg = GameMessage.deserializeBinary(event.data)
//     console.log(msg)
//     console.log(msg.type)
// })
