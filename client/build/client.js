"use strict";
exports.__esModule = true;
var net = require("net");
var conn = net.createConnection({
    host: "127.0.0.1",
    port: 8080
}, function () {
    conn.write("hello from the client");
    conn.write("yo");
    conn.write("my name is");
    conn.write(".smdjfjsldfm.");
});
conn.on("error", function (error) {
    console.log(error);
});
// const conn = new WebSocket("ws://localhost:8080")
// conn.onopen = () => {
//   conn.send("hello from the client!")
// }
// conn.onerror = (error) => {
//   console.log(error)
// }
