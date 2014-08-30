package channel

import "github.com/hawx/iirc/message"

type Client interface {
	Send(message.M)
	Name() string
}
