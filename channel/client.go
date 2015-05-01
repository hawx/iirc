package channel

import "hawx.me/code/iirc/message"

type Client interface {
	Send(message.M)
	Name() string
}
