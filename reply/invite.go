package reply

import "github.com/hawx/iirc/message"

func Invite(nick, user, host, subject, channel string) message.M {
	return message.Message3(
		message.Prefix(nick, user, host),
		"INVITE",
		message.ParamsT([]string{subject}, channel))
}
