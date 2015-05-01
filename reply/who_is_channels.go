package reply

import (
	"strings"

	"hawx.me/code/iirc/message"
)

// "<nick> :*( ( "@" / "+" ) <channel> " " )"
const RPL_WHOISCHANNELS = "319"

func WhoIsChannels(host, nick string, channels []string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISCHANNELS,
		message.ParamsT([]string{nick}, strings.Join(channels, " ")))
}
