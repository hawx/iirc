package reply

import "hawx.me/code/iirc/message"

const RPL_ENDOFWHO = "315"

func EndOfWho(host, user, subject string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_ENDOFWHO,
		message.Params([]string{user, subject}))
}
