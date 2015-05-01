package commands

import (
	"hawx.me/code/iirc/channel"
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Who(c Client, s Server, args []string) {
	if len(args) < 1 {
		c.Send(errors.NeedMoreParams(s.Name(), "WHO"))
		return
	}

	subject, ok := s.Find(args[0])

	if ok {
		if user, ok := subject.(Client); ok {
			c.Send(reply.WhoReply(s.Name(), c.Name(), args[0], user.UserName(), user.Name(), user.RealName()))
		} else {
			channel, _ := subject.(*channel.Channel)
			for _, client := range channel.Clients() {
				user, _ := client.(Client)
				c.Send(reply.WhoReply(s.Name(), c.Name(), args[0], user.UserName(), user.Name(), user.RealName()))
			}
		}
	}

	c.Send(reply.EndOfWho(s.Name(), c.Name(), args[0]))
}
