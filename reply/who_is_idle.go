package reply

import (
	"github.com/hawx/iirc/message"
	"strconv"
)

// "<nick> <integer> :seconds idle"
const RPL_WHOISIDLE = "317"

func WhoIsIdle(host, nick string, seconds int) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISIDLE,
		message.ParamsT([]string{nick, strconv.Itoa(seconds)}, "seconds idle"))
}
