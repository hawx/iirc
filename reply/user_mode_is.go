package reply

import "github.com/hawx/iirc/message"

 //"<user mode string>"
const RPL_UMODEIS = "221"

func UserModeIs(host, nick string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_UMODEIS,
		message.Params([]string{nick, "+"}))
}
