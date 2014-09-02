package handler

import "github.com/hawx/iirc/reply"

func Away(c Client, s Server, args []string) {
	if len(args) > 0 {
		c.SetAwayMessage(args[0])
		c.Send(reply.NowAway())
	} else {
		c.SetAwayMessage("")
		c.Send(reply.UnAway())
	}
}