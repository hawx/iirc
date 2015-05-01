package reply

import "hawx.me/code/iirc/message"

func NowAway(host string) message.M {
	return message.Message3(
		message.Prefix(host),
		"306",
		message.ParamsT([]string{}, "You have been marked as being away"))
}
