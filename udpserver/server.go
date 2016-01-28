package udpserver

import (
	"encoding/json"
	"fmt"
	"net"
)

// UDPServer is the server type
type UDPServer struct {
	Port string
}

// ClientMessage is a message sent from a client UDP connection
// to the server
type ClientMessage struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

// ToJSON converts a ClientMessage to json
func (c *ClientMessage) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// New creates a new UDPServer
func New(port string) *UDPServer {
	server := &UDPServer{
		Port: port,
	}
	return server
}

// Open opens the UDP server listener
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

// toClientMessage unmarshals byte[] to ClientMessage
func toClientMessage(byteArr []byte, bytesAvail int) (*ClientMessage, error) {
	var clientMsg *ClientMessage
	err := json.Unmarshal(byteArr[0:bytesAvail], &clientMsg)
	return clientMsg, err
}
