package reply

import "hawx.me/code/iirc/message"

func Nick(oldName, newName string) message.M {
	return message.Message3(
		message.Prefix(oldName),
		"NICK",
		message.Params([]string{newName}))
}
