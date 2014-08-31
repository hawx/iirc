package server

import (
	"container/list"
	"github.com/hawx/iirc/message"
)

func NewClients() *Clients {
	return &Clients{list.New()}
}

type Clients struct {
	l *list.List
}

func (cs *Clients) Len() int {
	return cs.l.Len()
}

func (cs *Clients) Add(c *Client) {
	cs.l.PushBack(c)
}

func (cs *Clients) Find(name string) (*Client, bool) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Client)
		if t.Name() == name {
			return t, true
		}
	}

	return nil, false
}

func (cs *Clients) Remove(c *Client) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Client)
		if t.Name() == c.Name() {
			cs.l.Remove(e)
		}
	}
}

func (cs *Clients) Broadcast(msg message.M) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		e.Value.(*Client).Send(msg)
	}
}

func (cs *Clients) Close() {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		e.Value.(*Client).Close()
	}
}

func (cs *Clients) Names() []string {
	names := []string{}
	for e := cs.l.Front(); e != nil; e = e.Next() {
		names = append(names, e.Value.(*Client).Name())
	}
	return names
}
