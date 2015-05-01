package commands

import "hawx.me/code/iirc/reply"

func Ping(c Client, s Server, args []string) {
	c.Send(reply.Pong(s.Name(), s.Address()))
}
