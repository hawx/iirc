package handler

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/message"
	"github.com/hawx/iirc/reply"
	"github.com/hawx/iirc/channel"
)

type Server2 interface {
	Address() string
	Name()    string
	Names()   []string
	FindChannel(string) *channel.Channel
	Join(channel.Client, string) *channel.Channel
	Part(channel.Client, string)
}

type Client2 interface {
	Send(message.M)
	Name() string
	SetName(string)
	UserName() string
	SetUserName(string)
	SetMode(string)
	RealName() string
	SetRealName(string)
	Channels() *channel.Channels
}

func Ping(c Client2, s Server2) {
	c.Send(reply.Pong(s.Address()))
}

func Nick(c Client2, s Server2, args []string) {
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

	oldName := c.Name()
	c.SetName(args[0])
	if oldName != "" {
		c.Send(reply.Nick(oldName, c.Name()))
	}
}

func User(c Client2, s Server2, args []string, command string) {
	if c.RealName() != "" {
		c.Send(errors.AlreadyRegistered())
		return
	}

	if len(args) < 4 {
		c.Send(errors.NeedMoreParams(command))
		return
	}

	// <user> <mode> <unused> <realname>
	c.SetUserName(args[0])
	c.SetMode(args[1])
	c.SetRealName(args[3])
	c.Send(reply.Welcome(s.Name(), c.Name()))
}

func Names(c Client2, s Server2, args []string) {
	channel := s.FindChannel(args[0])

	resp := reply.NameReply(s.Name(), c.Name(), channel.Name, channel.Names())

	for _, part := range resp.Parts() {
		c.Send(part)
	}

	c.Send(reply.EndOfNames(s.Name(), c.Name(), channel.Name))
}

func Join(c Client2, s Server2, args []string) {
	channel := s.Join(c, args[0])
	channel.Broadcast(reply.Join(c.Name(), c.UserName(), s.Name(), channel.Name))
	c.Channels().Add(channel)
}

func Part(c Client2, s Server2, args []string) {
	channel, ok := c.Channels().Find(args[0])

	if !ok {
		c.Send(errors.NotOnChannel(args[0]))
		return
	}

	channel.Broadcast(reply.Part(c.Name(), c.UserName(), s.Name(), channel.Name))
	s.Part(c, args[0])
}

func Topic(c Client2, s Server2, args []string) {
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
		channel.Broadcast(reply.TopicChange(c.Name(), c.UserName(), s.Name(), channel.Name, channel.Topic))

	default:
		return
	}
}
