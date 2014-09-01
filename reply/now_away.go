package reply

import "github.com/hawx/iirc/message"

func NowAway() message.M {
	return message.MessageParams(
		"306",
		message.ParamsT([]string{}, "You have been marked as being away"))
}
