package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

type UDPEchoServer struct {
	conn *net.UDPConn
}

// Creating the server struct
func NewUDPServer() UDPEchoServer {
	port, err := strconv.Atoi(PORT)
	addr := net.UDPAddr{IP: net.ParseIP(ADDRESS), Port: port}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal("Error creating server", err)
	}
	return UDPEchoServer{conn: conn}
}

// Listening for messages
func (server UDPEchoServer) Start() {
	msg := make([]byte, 64)
	for {
		n, addr, err := server.conn.ReadFromUDP(msg)
		log.Printf("Received message from %s", addr)
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client")
			} else {
				log.Println(err)
			}
			return
		}
		response := fmt.Sprintf("[ECHO-SERVER] %s", msg[:n])
		server.conn.WriteToUDP([]byte(response), addr)
	}
}
