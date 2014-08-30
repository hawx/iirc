package channel

import (
	"container/list"
	"github.com/hawx/iirc/message"
)

type Channel struct {
	Name    string
	Topic   string
	clients *list.List
}

func NewChannel(name string) *Channel {
	return &Channel{name, "", list.New()}
}

func (c *Channel) Join(client Client) {
	c.clients.PushBack(client)
}

func (c *Channel) Leave(client Client) {
	for e := c.clients.Front(); e != nil; e = e.Next() {
		t := e.Value.(Client)
		if t.Name() == client.Name() {
			c.clients.Remove(e)
		}
	}
}

func (c *Channel) Empty() bool {
	return c.clients.Len() == 0
}

func (c *Channel) Broadcast(msg message.M) {
	for e := c.clients.Front(); e != nil; e = e.Next() {
		e.Value.(Client).Send(msg)
	}
}

func (c *Channel) Names() []string {
	names := []string{}
	for e := c.clients.Front(); e != nil; e = e.Next() {
		names = append(names, e.Value.(Client).Name())
	}
	return names
}