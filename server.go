package gudp

import (
	"fmt"
	"net"
)

// Server receives the request from the Client
type Server struct {
	handlers
	Host       string
	Port       int
	BufferSize int
	err        error
}

// Listen opens up a connection and listens for messages
func (server *Server) Listen() (err error) {
	var address *net.UDPAddr
	var conn *net.UDPConn

	address, server.err = getAddress(server.Host, server.Port)
	conn, server.err = getConnection(address, "listen")

	defer conn.Close()

	buf := make([]byte, server.BufferSize)

	for {
		n, addr, err := conn.ReadFromUDP(buf)

		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	return
}
