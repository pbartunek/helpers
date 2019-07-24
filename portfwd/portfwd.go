package main

import (
  "io"
  "log"
  "net"
  "flag"
)

var (
  listenAddr string
  forwardAddr string
)

func forward(conn net.Conn) {
  client, err := net.Dial("tcp", forwardAddr)
  if err != nil {
    log.Printf("Connection failed: %v", err)
    log.Printf("Closing client connection")
    conn.Close()
    return
  }

  log.Printf("Connected to %v\n", client.RemoteAddr())
  go func() {
    defer client.Close()
    defer conn.Close()
    io.Copy(client, conn)
  }()

  go func() {
    defer client.Close()
    defer conn.Close()
    io.Copy(conn, client)
  }()
}

func main() {
  flag.StringVar(&listenAddr, "l", "127.0.0.1:1337", "listen ip:port")
  flag.StringVar(&forwardAddr, "f", "127.0.0.1:8080", "forward ip:port")
  flag.Parse()

  log.Printf("Starting port forwarder, listen address: %v, forwarding to: %v\n", listenAddr, forwardAddr)
  listener, err := net.Listen("tcp", listenAddr)
  if err != nil {
    log.Fatalf("Failed to setup listener: %v", err)
  }

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Fatalf("Failed to accept listener: %v", err)
    }
    log.Printf("Accepted connection %v\n", conn.RemoteAddr())
    go forward(conn)
  }
}
