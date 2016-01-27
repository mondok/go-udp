package udpserver

import (
	"fmt"
	"net"
)

type UDPServer struct {
	Port string
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
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			return err
		}
	}
}
