package reply

import "github.com/hawx/iirc/message"

func UnAway() message.M {
	return message.MessageParams(
		"305",
		message.ParamsT([]string{}, "You are no longer marked as being away"))
}
