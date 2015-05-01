package channel

import (
	"container/list"

	"hawx.me/code/iirc/message"
)

type Channel struct {
	name    string
	Topic   string
	clients *list.List
}

func NewChannel(name string) *Channel {
	return &Channel{name, "", list.New()}
}

func (c *Channel) Name() string {
	return c.name
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

func (c *Channel) Clients() []Client {
	clients := []Client{}

	for e := c.clients.Front(); e != nil; e = e.Next() {
		clients = append(clients, e.Value.(Client))
	}

	return clients
}

func (c *Channel) Empty() bool {
	return c.clients.Len() == 0
}

func (c *Channel) Send(msg message.M) {
	for e := c.clients.Front(); e != nil; e = e.Next() {
		e.Value.(Client).Send(msg)
	}
}

func (c *Channel) SendExcept(msg message.M, name string) {
	for e := c.clients.Front(); e != nil; e = e.Next() {
		t := e.Value.(Client)
		if t.Name() != name {
			t.Send(msg)
		}
	}
}

func (c *Channel) Names() []string {
	names := []string{}
	for e := c.clients.Front(); e != nil; e = e.Next() {
		names = append(names, e.Value.(Client).Name())
	}
	return names
}
