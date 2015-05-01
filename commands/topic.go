package commands

import (
	"hawx.me/code/iirc/errors"
	"hawx.me/code/iirc/reply"
)

func Topic(c Client, s Server, args []string) {
	switch len(args) {
	case 1:
		channel := s.FindChannel(args[0])
		if channel.Topic == "" {
			c.Send(reply.NoTopic(s.Name(), channel.Name()))
		} else {
			c.Send(reply.Topic(s.Name(), channel.Name(), channel.Topic))
		}

	case 2:
		channel := s.FindChannel(args[0])
		channel.Topic = args[1]
		channel.Send(reply.TopicChange(c.Name(), c.UserName(), s.Name(), channel.Name(), channel.Topic))

	default:
		c.Send(errors.NeedMoreParams(s.Name(), "TOPIC"))
	}
}
