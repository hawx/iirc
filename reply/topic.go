package reply

import "hawx.me/code/iirc/message"

// "<channel> :<topic>"
const RPL_TOPIC = "332"

func Topic(host, channel, topic string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_TOPIC,
		message.ParamsT([]string{channel}, topic))
}
