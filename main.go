package main

import (
	"log"
)

func main() {
	echoSRV := NewTCPServer()
	defer echoSRV.listener.Close()
	for {
		conn, err := echoSRV.listener.Accept()
		if err != nil {
			log.Printf("Could not accept connection %s - %s", conn.RemoteAddr(), err)
			continue
		}
		go echoSRV.HandleConn(conn)
	}
}
