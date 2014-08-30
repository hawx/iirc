package reply

import "github.com/hawx/iirc/message"

func Pong(address string) message.M {
	return message.MessageParams("PONG", message.Params([]string{address}))
}
