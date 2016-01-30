package main

import (
	"fmt"
	"net"
	"time"

	"github.com/mondok/go-udp/udpserver"
)

func main() {
	server := udpserver.New("10001")
	go server.Open()
	exampleClient("127.0.0.1:10001")
}

func exampleClient(address string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		_ = fmt.Errorf("Error %v", err)
		return
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		_ = fmt.Errorf("Error %v", err)
		return
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		_ = fmt.Errorf("Error %v", err)
		return
	}

	defer Conn.Close()
	for {
		msg := &udpserver.ClientMessage{}
		msg.ID = "123"
		msg.Body = "Hello World"
		buf, _ := msg.ToJSON()
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
