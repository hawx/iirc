package handler

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func User(c Client, s Server, args []string, command string) {
	if c.RealName() != "" {
		c.Send(errors.AlreadyRegistered())
		return
	}

	if len(args) < 4 {
		c.Send(errors.NeedMoreParams(command))
		return
	}

	// <user> <mode> <unused> <realname>
	c.SetUserName(args[0])
	c.SetMode(args[1])
	c.SetRealName(args[3])
	c.Send(reply.Welcome(s.Name(), c.Name()))
}
