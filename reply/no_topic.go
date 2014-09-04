package reply

import "github.com/hawx/iirc/message"

// "<channel> :No topic is set"
const RPL_NOTOPIC = "331"

func NoTopic(host, channel string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_NOTOPIC,
		message.ParamsT([]string{channel}, "No topic is set"))
}
