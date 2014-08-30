package reply

import (
	"github.com/hawx/iirc/message"
	"strings"
)

//   "( "=" / "*" / "@" ) <channel> :[ "@" / "+" ] <nick> *( " " [ "@" / "+" ] <nick> )
// "@" is used for secret channels, "*" for private channels, and "=" for
// others (public channels).
const	RPL_NAMREPLY = "353"

func NameReply(serverName, nickName, channelName string, names []string) message.M {
	return message.Message3(
		message.Prefix(serverName),
		RPL_NAMREPLY,
		message.ParamsT([]string{nickName, "=", channelName}, strings.Join(names, " ")))
}
