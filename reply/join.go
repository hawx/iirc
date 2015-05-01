package reply

import "hawx.me/code/iirc/message"

func Join(nickName, userName, serverName, channelName string) message.M {
	return message.Message3(
		message.Prefix(nickName, userName, serverName),
		"JOIN",
		message.ParamsT([]string{}, channelName))
}
