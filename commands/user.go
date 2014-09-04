package commands

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func User(c Client, s Server, args []string) {
	if c.RealName() != "" {
		c.Send(errors.AlreadyRegistered(s.Name()))
		return
	}

	if len(args) < 4 {
		c.Send(errors.NeedMoreParams(s.Name(), "USER"))
		return
	}

	// <user> <mode> <unused> <realname>
	c.SetUserName(args[0])
	c.SetMode(args[1])
	c.SetRealName(args[3])
	c.Send(reply.Welcome(s.Name(), c.Name()))
}
