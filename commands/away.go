package commands

import "github.com/hawx/iirc/reply"

func Away(c Client, s Server, args []string) {
	if len(args) > 0 {
		c.SetAwayMessage(args[0])
		c.Send(reply.NowAway(s.Name()))
		return
	}

	c.SetAwayMessage("")
	c.Send(reply.UnAway(s.Name()))
}
