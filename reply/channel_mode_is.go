package reply

import "github.com/hawx/iirc/message"

// "<channel> <mode> <mode params>"
const RPL_CHANNELMODEIS = "324"

func ChannelModeIs(host, nick, channel string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_CHANNELMODEIS,
		message.Params([]string{nick, channel, "+"}))
}
