package reply

import "github.com/hawx/iirc/message"

// "<channel> :<topic>"
const RPL_TOPIC = "332"

func Topic(channel, topic string) message.M {
	return message.MessageParams(
		RPL_TOPIC,
		message.ParamsT([]string{channel}, topic))
}
