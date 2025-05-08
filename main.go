package main

const (
	PORT    = "3000"
	ADDRESS = "localhost"
)

func main() {
	echoUDPServer := NewUDPServer()
	echoTCPServer := NewTCPServer()
	go echoUDPServer.Start()
	echoTCPServer.Start()
}
