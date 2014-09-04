package commands

import "github.com/hawx/iirc/reply"

func Ping(c Client, s Server, args []string) {
	c.Send(reply.Pong(s.Name(), s.Address()))
}
