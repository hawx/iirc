package reply

import "github.com/hawx/iirc/message"

const RPL_ENDOFWHO = "315"

func EndOfWho(host, user, subject string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_ENDOFWHO,
		message.Params([]string{user, subject}))
}
