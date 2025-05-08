package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	PROTOCOL = "tcp"
	PORT     = "3000"
)

type EchoServer struct {
	listener net.Listener
}

func NewServer() EchoServer {
	l, err := net.Listen(PROTOCOL, ":"+PORT)
	if err != nil {
		log.Fatal("Error creating server", err)
	}
	return EchoServer{listener: l}
}
func (server EchoServer) HandleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("Could not read any more messages from client", err)
			return
		}
		response := fmt.Sprintf("[ECHO-SERVER] %s", msg)
		conn.Write([]byte(response))
	}

}
