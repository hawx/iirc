package commands

import (
	"github.com/hawx/iirc/reply"
	"github.com/hawx/iirc/errors"
)


func Whois(c Client, s Server, args []string) {
	subject, ok := s.Find(args[0])
	user, isClient := subject.(Client)

	if !ok || !isClient {
		c.Send(errors.NoSuchNick(s.Name(), args[0]))
		c.Send(reply.EndOfWhois(s.Name(), args[0]))
		return
	}

	if user.AwayMessage() != "" {
		c.Send(reply.Away(s.Name(), user.Name(), user.AwayMessage()))
	}

	c.Send(reply.WhoIsUser(s.Name(), user.Name(), user.UserName(), user.RealName()))
	c.Send(reply.WhoIsServer(s.Name(), user.Name()))

	if user.Channels().Any() {
		resp := reply.WhoIsChannels(s.Name(), user.Name(), user.Channels().Names())
		for _, part := range resp.Parts() {
			c.Send(part)
		}
	}

	c.Send(reply.WhoIsIdle(s.Name(), user.Name(), 0))
	c.Send(reply.EndOfWhois(s.Name(), args[0]))
}
