package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func User(c Client, s Server, args []string) {
	if len(args) < 4 {
		c.Send(errors.NeedMoreParams(s.Name(), "USER"))
		return
	}

	if c.RealName() != "" {
		c.Send(errors.AlreadyRegistered(s.Name()))
		return
	}

	// <user> <mode> <unused> <realname>
	c.SetUserName(args[0])
	c.SetMode(args[1])
	c.SetRealName(args[3])
	c.Send(reply.Welcome(s.Name(), c.Name()))
}
