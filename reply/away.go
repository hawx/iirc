package reply

import "hawx.me/code/iirc/message"

// "<nick> :<away message>"
const RPL_AWAY = "301"

func Away(host, nick, awayMsg string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_AWAY,
		message.ParamsT([]string{nick}, awayMsg))
}
