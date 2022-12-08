// import { WebSocket } from "ws"
import * as net from "net"

const conn = net.createConnection({
  host: "127.0.0.1",
  port: 8080
}, () => {
  conn.write("hello from the client")
})

conn.on("error", (error) => {
  console.log(error)
})

// const conn = new WebSocket("ws://localhost:8080")

// conn.onopen = () => {
//   conn.send("hello from the client!")
// }

// conn.onerror = (error) => {
//   console.log(error)
// }
