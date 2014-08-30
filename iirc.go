package main

import (
	"github.com/hawx/iirc/server"
)

func main() {
	const (
		address    = "127.0.0.1"
		port       = "6767"
		serverName = "hawx.irc"
	)

	s := server.NewServer(serverName, address, port)
	s.Start()
}
