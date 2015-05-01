package connection

import "hawx.me/code/iirc/message"

type Handler interface {
	OnSend(message.M)
	OnReceive(message.M)
	OnError(error)
	OnQuit()
}
