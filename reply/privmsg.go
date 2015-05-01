package reply

import "hawx.me/code/iirc/message"

func PrivMsg(nick, user, host, subject, msg string) message.M {
	return message.Message3(
		message.Prefix(nick, user, host),
		"PRIVMSG",
		message.ParamsT([]string{subject}, msg))
}
