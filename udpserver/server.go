package udpserver

import (
	"fmt"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/mondok/go-udp/udpserver"
)

// UDPServer is the server type
type UDPServer struct {
	Port string
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
func toClientMessage(byteArr []byte, bytesAvail int) (*udpserver.ClientMessage, error) {
	var clientMsg *udpserver.ClientMessage
	proto.Unmarshal(byteArr, &clientMsg)
	return clientMsg, err
}
