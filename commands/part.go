package commands

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func Part(c Client, s Server, args []string) {
	channel, ok := c.Channels().Find(args[0])

	if !ok {
		c.Send(errors.NotOnChannel(s.Name(), args[0]))
		return
	}

	channel.Send(reply.Part(c.Name(), c.UserName(), s.Name(), channel.Name()))
	s.Part(c, args[0])
}
