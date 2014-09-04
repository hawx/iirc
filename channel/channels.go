package channel

import (
	"container/list"
)

func NewChannels() *Channels {
	return &Channels{list.New()}
}

type Channels struct {
	l *list.List
}

func (cs *Channels) Any() bool {
	return cs.l.Len() > 0
}

func (cs *Channels) Each(f func(*Channel)) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		f(e.Value.(*Channel))
	}
}

func (cs *Channels) Find(name string) (*Channel, bool) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Channel)
		if t.Name() == name {
			return t, true
		}
	}

	return NewChannel(name), false
}

func (cs *Channels) Add(c *Channel) {
	cs.l.PushBack(c)
}

func (cs *Channels) Remove(c *Channel) {
	for e := cs.l.Front(); e != nil; e = e.Next() {
		t := e.Value.(*Channel)
		if t.Name() == c.Name() {
			cs.l.Remove(e)
		}
	}
}

func (cs *Channels) Names() []string {
	names := []string{}
	cs.Each(func(c *Channel) {
		names = append(names, c.Name())
	})

	return names
}
