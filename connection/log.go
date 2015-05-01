package connection

import (
	"log"

	"hawx.me/code/iirc/message"
)

type logHandler struct {
	name  string
	inner Handler
}

func Log(name string, inner Handler) Handler {
	return logHandler{name, inner}
}

func (h logHandler) OnSend(m message.M) {
	log.Println(h.name, " <- ", m)
	h.inner.OnSend(m)
}

func (h logHandler) OnReceive(m message.M) {
	log.Print(h.name, " -> ", m)
	h.inner.OnReceive(m)
}

func (h logHandler) OnError(e error) {
	log.Println(e)
	h.inner.OnError(e)
}

func (h logHandler) OnQuit() {
	h.inner.OnQuit()
}
