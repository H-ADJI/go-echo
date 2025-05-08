package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type TCPEchoServer struct {
	listener net.Listener
}

func NewTCPServer() TCPEchoServer {
	l, err := net.Listen("tcp", ADDRESS+":"+PORT)
	if err != nil {
		log.Fatal("Error creating server", err)
	}
	return TCPEchoServer{listener: l}
}

// Starts listening
func (server TCPEchoServer) Start() {
	defer server.listener.Close()
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection %s - %s", conn.RemoteAddr(), err)
			continue
		}
		go server.HandleConn(conn)
	}
}

// Handle each TCP connection
func (server TCPEchoServer) HandleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client")
			} else {
				log.Println(err)
			}
			return
		}
		response := fmt.Sprintf("[ECHO-SERVER] %s", msg)
		conn.Write([]byte(response))
	}

}
