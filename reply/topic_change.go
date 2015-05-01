package reply

import "hawx.me/code/iirc/message"

func TopicChange(nickName, userName, serverName, channel, topic string) message.M {
	return message.Message3(
		message.Prefix(nickName, userName, serverName),
		"TOPIC",
		message.ParamsT([]string{channel}, topic))
}
