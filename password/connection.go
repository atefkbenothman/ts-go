package app

import (
  "fmt"
  "log"
  "net"
)

func ReadFromConnection(host string) {
  // create TCP server
  server, err := net.Listen("tcp", host)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("started server @ %s\n", host)
  fmt.Println(server.Addr())

  // accept incoming connection
  conn, err := server.Accept()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("accepted connection")

  // defer: close the connection when the function returns
  defer conn.Close()

  // create byte slice
  buf := make([]byte, 1024)

  for {
    numBytes, err := conn.Read(buf)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(string(buf[:numBytes]))
  }
}
