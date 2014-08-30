package reply

import "github.com/hawx/iirc/message"

func Nick(oldName, newName string) message.M {
	return message.Message3(
		message.Prefix(oldName),
		"NICK",
		message.Params([]string{newName}))
}
