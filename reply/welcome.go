package reply

import "hawx.me/code/iirc/message"

// "Welcome to the Internet Relay Network <nick>!<user>@<host>"
const RPL_WELCOME = "001"

func Welcome(serverName, nickName string) message.M {
	return message.Message3(
		message.Prefix(serverName),
		RPL_WELCOME,
		message.Params([]string{nickName}))
}
