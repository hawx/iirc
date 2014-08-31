package handler

import "github.com/hawx/iirc/reply"

func Topic(c Client, s Server, args []string) {
	switch len(args) {
	case 1:
		channel := s.FindChannel(args[0])
		if channel.Topic == "" {
			c.Send(reply.NoTopic(channel.Name()))
		} else {
			c.Send(reply.Topic(channel.Name(), channel.Topic))
		}

	case 2:
		channel := s.FindChannel(args[0])
		channel.Topic = args[1]
		channel.Send(reply.TopicChange(c.Name(), c.UserName(), s.Name(), channel.Name(), channel.Topic))

	default:
		return
	}
}
