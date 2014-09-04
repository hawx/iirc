package reply

import "github.com/hawx/iirc/message"

func UnAway(host string) message.M {
	return message.Message3(
		message.Prefix(host),
		"305",
		message.ParamsT([]string{}, "You are no longer marked as being away"))
}
