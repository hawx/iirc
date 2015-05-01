package commands

import (
	"strings"

	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Join(c Client, s Server, args []string) {
	if len(args) < 1 {
		c.Send(errors.NeedMoreParams(s.Name(), "JOIN"))
		return
	}

	names := strings.Split(args[0], ",")

	for _, name := range names {
		channel := s.Join(c, name)
		channel.Send(reply.Join(c.Name(), c.UserName(), s.Name(), channel.Name()))
		c.Channels().Add(channel)
		Topic(c, s, []string{name})
	}
}
