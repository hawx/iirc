package reply

import "github.com/hawx/iirc/message"

// "<channel> :End of channel ban list"
const RPL_ENDOFBANLIST = "368"

func EndOfBanList(host, nick, channel string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_ENDOFBANLIST,
		message.ParamsT([]string{nick, channel}, "End of Channel Ban List"))
}
