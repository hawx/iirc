package reply

import "github.com/hawx/iirc/message"

// "<channel> :No topic is set"
const RPL_NOTOPIC = "331"

func NoTopic(channel string) message.M {
	return message.MessageParams(
		RPL_NOTOPIC,
		message.ParamsT([]string{channel}, "No topic is set"))
}
