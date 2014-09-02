package handler

import (
	"github.com/hawx/iirc/reply"
	"strings"
)

func Mode(c Client, s Server, args []string) {
	subject, ok := s.Find(args[0])

	if !ok { return }

	if _, ok := subject.(Client); ok {
		c.Send(reply.UserModeIs(s.Name(), c.Name()))
	} else {
		if len(args) == 2 && strings.Contains(args[1], "b") {
			c.Send(reply.EndOfBanList(s.Name(), c.Name(), subject.Name()))
		} else {
			c.Send(reply.ChannelModeIs(s.Name(), c.Name(), subject.Name()))
		}
	}
}
