package reply

import "github.com/hawx/iirc/message"

// "<channel> <user> <host> <server> <nick>
//   ( "H" / "G" > ["*"] [ ( "@" / "+" )]
//   :<hopcount> <real name>"
const RPL_WHOREPLY = "352"

func WhoReply(host, me, subject, user, nick, realName string) message.M {
	return message.Message3(
		message.Prefix(host),
		RPL_WHOREPLY,
		message.ParamsT([]string{me, subject, user, host, host, nick, "H"}, "0 " + realName))
}
