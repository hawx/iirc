package handler

import (
	"github.com/hawx/iirc/errors"
	"github.com/hawx/iirc/reply"
)

func Nick(c Client, s Server, args []string) {
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