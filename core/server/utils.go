package server

import (
	"net"
)

func isPortInUse(port string) bool {
	listener, err := net.Listen("tcp", net.JoinHostPort("", port))

	if err != nil {
		return true
	}

	defer listener.Close()
	return false
}
