package commands

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func Names(c Client, s Server, args []string) {
	if len(args) < 1 {
		c.Send(errors.NeedMoreParams(s.Name(), "NAMES"))
		return
	}

	channel := s.FindChannel(args[0])

	resp := reply.NameReply(s.Name(), c.Name(), channel.Name(), channel.Names())

	for _, part := range resp.Parts() {
		c.Send(part)
	}

	c.Send(reply.EndOfNames(s.Name(), c.Name(), channel.Name()))
}
