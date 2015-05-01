package reply

import "hawx.me/code/iirc/message"

func Part(nickName, userName, serverName, channel string) message.M {
	return message.Message3(
		message.Prefix(nickName, userName, serverName),
		"PART",
		message.ParamsT([]string{channel}, nickName))
}
