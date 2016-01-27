package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/mondok/goudp/udpserver"
)

func main() {
	server := udpserver.New("10001")
	go server.Open()
	exampleClient("127.0.0.1:10001")
}

func exampleClient(address string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		fmt.Errorf("Error %v", err)
		return
	}

	defer Conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}
