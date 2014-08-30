package reply

import "github.com/hawx/iirc/message"

// "<channel> :End of NAMES list"
const RPL_ENDOFNAMES = "366"

func EndOfNames(serverName, nickName, channelName string) message.M {
	return message.Message3(
		message.Prefix(serverName),
		RPL_ENDOFNAMES,
		message.ParamsT([]string{nickName, channelName}, "End of /NAMES list."))
}
