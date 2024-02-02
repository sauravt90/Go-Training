package main

import (
	"fmt"
)

const (
	serverIP = "0.0.0.0"
	port     = "1234"
)

func main() {
	fmt.Println("Demo TCP server")

	address := serverIP + ":" + port

	fmt.Println("Listning on at address", address)

	//net.Listen()
}
