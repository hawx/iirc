package reply

import "github.com/hawx/iirc/message"

// "<nick> :is an IRC operator"
const RPL_WHOISOPERATOR = "313"

func WhoIsOperator(host, nick string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOISOPERATOR,
		message.ParamsT([]string{nick}, "is an IRC oeprator"))
}
