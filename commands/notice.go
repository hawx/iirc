package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Notice(c Client, s Server, args []string) {
	if len(args) < 2 {
		c.Send(errors.NeedMoreParams(s.Name(), "NOTICE"))
		return
	}

	subject, ok := s.Find(args[0])
	if ok {
		msg := reply.Notice(c.Name(), c.UserName(), s.Name(), args[0], args[1])
		subject.SendExcept(msg, c.Name())
	}
}
