package main

import (
	"github.com/hawx/iirc/server"
	"flag"
)

var (
	addr = flag.String("addr", "127.0.0.1", "")
	port = flag.String("port", "6767", "")
	name = flag.String("name", "iirc.irc", "")
)

func main() {
	flag.Parse()
	s := server.NewServer(*name, *addr, *port)
	s.Start()
}
