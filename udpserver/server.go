package udpserver

import (
	"encoding/json"
	"fmt"
	"net"
)

type UDPServer struct {
	Port string
}

type ClientMessage struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

func (c *ClientMessage) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

func New(port string) *UDPServer {
	server := &UDPServer{
		Port: port,
	}
	return server
}

func (u *UDPServer) Open() error {
	ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", u.Port))
	if err != nil {
		return err
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		return err
	}
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		msg, _ := toClientMessage(buf, n)
		fmt.Println("Message ", msg.Body, " from ", addr)
		if err != nil {
			return err
		}
	}
}

func toClientMessage(byteArr []byte, bytesAvail int) (*ClientMessage, error) {
	var clientMsg *ClientMessage
	err := json.Unmarshal(byteArr[0:bytesAvail], &clientMsg)
	return clientMsg, err
}
