package reply

import "hawx.me/code/iirc/message"

func Pong(host, address string) message.M {
	return message.Message3(
		message.Prefix(host),
		"PONG",
		message.Params([]string{address}))
}
