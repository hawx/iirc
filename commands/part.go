package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Part(c Client, s Server, args []string) {
	if len(args) < 1 {
		c.Send(errors.NeedMoreParams(s.Name(), "PART"))
		return
	}

	channel, ok := c.Channels().Find(args[0])

	if !ok {
		c.Send(errors.NotOnChannel(s.Name(), args[0]))
		return
	}

	channel.Send(reply.Part(c.Name(), c.UserName(), s.Name(), channel.Name()))
	s.Part(c, args[0])
}
