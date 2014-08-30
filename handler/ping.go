package handler

import "github.com/hawx/iirc/reply"

func Ping(c Client, s Server) {
	c.Send(reply.Pong(s.Address()))
}
