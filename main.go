package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello world")
	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("Error creating server", err)
	}

	for {
		conn, err := l.Accept()
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			log.Fatal("Error while accepting connection", err)
		}
		go HandleConnection(conn)
	}

}
func HandleConnection(conn net.Conn) {
	message := make([]byte, 0)
	_, err := conn.Read(message)
	if err != nil {
		log.Fatal("error while reading from connection")
	}
	fmt.Printf("%s \n", message)
}
