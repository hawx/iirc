package commands

import "github.com/hawx/iirc/reply"

func Names(c Client, s Server, args []string) {
	channel := s.FindChannel(args[0])

	resp := reply.NameReply(s.Name(), c.Name(), channel.Name(), channel.Names())

	for _, part := range resp.Parts() {
		c.Send(part)
	}

	c.Send(reply.EndOfNames(s.Name(), c.Name(), channel.Name()))
}
