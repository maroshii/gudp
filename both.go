package gudp

import (
	"errors"
	"fmt"
	"net"
)

type handlers interface {
	getConnection(*net.UDPAddr) (*net.UDPConn, error)
	getAddress(string, int) (*net.UDPAddr, error)
}

// getConnection creates the UDP connection for either side (Client or Server)
func getConnection(address *net.UDPAddr, action string) (conn *net.UDPConn, err error) {
	if action == "listen" {
		conn, err = net.ListenUDP("udp", address)
	} else if action == "dial" {
		conn, err = net.DialUDP("udp", nil, address)
	} else {
		err = errors.New("Invalid action. Available options are 'listen' and 'dial'")
	}

	return conn, err
}

// getAddress is a wrapper for the net.ResolveUDPAddr method
// It resolves the full server address given a host and a port
func getAddress(host string, port int) (udpAddr *net.UDPAddr, err error) {
	fullAddr := fmt.Sprintf("%s:%d", host, port)
	udpAddr, err = net.ResolveUDPAddr("udp", fullAddr)
	return udpAddr, err
}
