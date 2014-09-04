package reply

import (
	"github.com/hawx/iirc/message"
	"strings"
)

// "<nick> :*( ( "@" / "+" ) <channel> " " )"
const RPL_WHOISCHANNELS = "319"

func WhoIsChannels(host, nick string, channels []string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISCHANNELS,
		message.ParamsT([]string{nick}, strings.Join(channels, " ")))
}
