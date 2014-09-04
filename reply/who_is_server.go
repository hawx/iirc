package reply

import "github.com/hawx/iirc/message"

// "<nick> <server> :<server info>"
const RPL_WHOISSERVER = "312"

func WhoIsServer(host, user string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISSERVER,
		message.ParamsT([]string{user, host}, ""))
}
