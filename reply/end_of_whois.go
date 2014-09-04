package reply

import "github.com/hawx/iirc/message"

// "<nick> :End of WHOIS list"
const RPL_ENDOFWHOIS = "318"

func EndOfWhois(host, subject string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_ENDOFWHOIS,
		message.ParamsT([]string{subject}, "End of WHOIS list"))
}
