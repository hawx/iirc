package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func PrivMsg(c Client, s Server, args []string) {
	if len(args) < 2 {
		c.Send(errors.NeedMoreParams(s.Name(), "PRIVMSG"))
		return
	}

	subject, ok := s.Find(args[0])
	if ok {
		msg := reply.PrivMsg(c.Name(), c.UserName(), s.Name(), args[0], args[1])
		subject.SendExcept(msg, c.Name())

		if other, ok := subject.(Client); ok && other.AwayMessage() != "" {
			c.Send(reply.Away(s.Name(), subject.Name(), other.AwayMessage()))
		}
	}
}
