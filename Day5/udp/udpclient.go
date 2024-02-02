package main

import (
	"fmt"
	"net"
	"os"
)

const (
	serverIP = "0.0.0.0"
	port     = "2345"
)

func main() {
	fmt.Println("UDP client.........")

	address := serverIP + ":" + port

	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println("Recieved error ")
		os.Exit(1)
	}
	var messages = []string{"HELLO", "1234", "abs", ``}

	for _, msg := range messages {
		n, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("error iiw")
			continue
		}
		fmt.Println("Wrote ", n, "bytes")
		data := make([]byte, 1024)
		n, err = conn.Read(data)
		if err != nil {
			fmt.Println("error whil reading")
			continue
		}
		fmt.Println("Read ", n, "Bytes", string(data))
	}

}
