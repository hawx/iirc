package connection

import "github.com/hawx/iirc/message"

type Handler interface {
	OnSend(message.M)
	OnReceive(message.M)
	OnError(error)
	OnQuit()
}
