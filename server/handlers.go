package server

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func Ping(c *Client, s *Server) {
	c.Send(reply.Pong(s.Address()))
}

func Nick(c *Client, s *Server, args []string) {
	if len(args) == 0 {
		c.Send(errors.NoNicknameGiven())
		return
	}

	for _, name := range s.Names() {
		if name == args[0] {
			c.Send(errors.NicknameInUse(args[0]))
			return
		}
	}

	oldName := c.Name
	c.Name = args[0]
	if oldName != "" {
		c.Send(reply.Nick(oldName, c.Name))
	}
}

func User(c *Client, s *Server, args []string, command string) {
	if c.realName != "" {
		c.Send(errors.AlreadyRegistered())
		return
	}

	if len(args) < 4 {
		c.Send(errors.NeedMoreParams(command))
		return
	}

	// <user> <mode> <unused> <realname>
	c.userName = args[0]
	c.mode = args[1]
	c.realName = args[3]
	c.Send(reply.Welcome(s.Name(), c.Name))
}

func Names(c *Client, s *Server, args []string) {
	channel := s.FindChannel(args[0])

	resp := reply.NameReply(s.Name(), c.Name, channel.Name, channel.Names())

	for _, part := range resp.Parts() {
		c.Send(part)
	}

	c.Send(reply.EndOfNames(s.Name(), c.Name, channel.Name))
}

func Join(c *Client, s *Server, args []string) {
	channel := s.Join(c, args[0])
	channel.Broadcast(reply.Join(c.Name, c.userName, s.Name(), channel.Name))
	c.channels.Add(channel)
}

func Part(c *Client, s *Server, args []string) {
	channel, ok := c.channels.Find(args[0])

	if !ok {
		c.Send(errors.NotOnChannel(args[0]))
		return
	}

	channel.Broadcast(reply.Part(c.Name, c.userName, s.Name(), channel.Name))
	s.Part(c, args[0])
}

func Topic(c *Client, s *Server, args []string) {
	switch len(args) {
	case 1:
		channel := s.FindChannel(args[0])
		if channel.Topic == "" {
			c.Send(reply.NoTopic(channel.Name))
		} else {
			c.Send(reply.Topic(channel.Name, channel.Topic))
		}

	case 2:
		channel := s.FindChannel(args[0])
		channel.Topic = args[1]
		channel.Broadcast(reply.TopicChange(c.Name, c.userName, s.Name(), channel.Name, channel.Topic))

	default:
		return
	}
}
