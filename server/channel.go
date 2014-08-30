package server

import (
	"github.com/hawx/iirc/message"
)

type Channel struct {
	Name    string
	Topic   string
	clients *Clients
}

func NewChannel(name string) *Channel {
	return &Channel{name, "", NewClients()}
}

func (ch *Channel) Join(c *Client) {
	ch.clients.Add(c)
}

func (ch *Channel) Leave(c *Client) {
	ch.clients.Remove(c)
}

func (ch *Channel) Empty() bool {
	return ch.clients.Len() == 0
}

func (ch *Channel) Broadcast(msg message.M) {
	ch.clients.Broadcast(msg)
}

func (ch *Channel) Names() []string {
	return ch.clients.Names()
}
