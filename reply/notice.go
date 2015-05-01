package reply

import "hawx.me/code/iirc/message"

func Notice(nick, user, host, subject, msg string) message.M {
	return message.Message3(
		message.Prefix(nick, user, host),
		"NOTICE",
		message.ParamsT([]string{subject}, msg))
}
