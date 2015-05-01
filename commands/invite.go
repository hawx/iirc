package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Invite(c Client, s Server, args []string) {
	if len(args) < 2 {
		c.Send(errors.NeedMoreParams(s.Name(), "INVITE"))
		return
	}

	subject, ok := s.Find(args[0])
	user, isClient := subject.(Client)

	if !ok || !isClient {
		c.Send(errors.NoSuchNick(s.Name(), args[0]))
		return
	}

	if _, ok := c.Channels().Find(args[1]); !ok {
		c.Send(errors.NotOnChannel(s.Name(), args[1]))
		return
	}

	if _, ok := user.Channels().Find(args[1]); ok {
		c.Send(errors.UserOnChannel(s.Name(), user.Name(), args[1]))
		return
	}

	subject.Send(reply.Invite(c.Name(), c.UserName(), s.Name(), subject.Name(), args[1]))
}
