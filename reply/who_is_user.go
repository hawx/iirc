package reply

import "github.com/hawx/iirc/message"

// "<nick> <user> <host> * :<real name>"
const RPL_WHOISUSER = "311"

func WhoIsUser(host, nick, user, realName string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISUSER,
		message.ParamsT([]string{nick, user, host, "*"}, realName))
}
