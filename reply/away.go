package reply

import "github.com/hawx/iirc/message"

func Away(nick, awayMsg string) message.M {
	return message.MessageParams(
		"301",
		message.ParamsT([]string{nick}, awayMsg))
}
