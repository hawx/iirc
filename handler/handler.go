package handler

import (
	"github.com/hawx/iirc/channel"
	"github.com/hawx/iirc/message"
)

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
