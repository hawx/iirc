package commands

import (
	"hawx.me/code/iirc/channel"
	"hawx.me/code/iirc/message"
)

type Command func(Client, Server, []string)

type Server interface {
	Address() string
	Name() string
	Names() []string
	FindChannel(string) *channel.Channel
	Find(string) (message.Sender, bool)
	Join(channel.Client, string) *channel.Channel
	Part(channel.Client, string)
}

type Client interface {
	Send(message.M)
	Name() string
	SetName(string)
	UserName() string
	SetUserName(string)
	SetMode(string)
	RealName() string
	SetRealName(string)
	Channels() *channel.Channels
	AwayMessage() string
	SetAwayMessage(string)
}
