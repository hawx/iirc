package handler

import "github.com/hawx/iirc/reply"

func Join(c Client, s Server, args []string) {
	channel := s.Join(c, args[0])
	channel.Broadcast(reply.Join(c.Name(), c.UserName(), s.Name(), channel.Name))
	c.Channels().Add(channel)
}
