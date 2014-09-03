package connection

import (
	"github.com/hawx/iirc/message"
	"net"
	"bufio"
)

type Conn interface {
	Send(message.M)
	Close()
}

type conn struct {
	in      chan message.M
	out     chan string
	quit    chan struct{}
	conn    net.Conn
	handler Handler
}

func NewConn(netConn net.Conn, handler Handler) *conn {
	c := &conn{
	  in: make(chan message.M),
	  out: make(chan string),
  	quit: make(chan struct{}),
  	conn: netConn,
  	handler: handler,
	}

	go c.receiver()
	go c.sender()

	return c
}

func (c *conn) Send(msg message.M) {
	c.in <- msg
}

func (c *conn) Close() {
	c.quit <- struct{}{}
}

func (c *conn) receiver() {
	r := bufio.NewReader(c.conn)

	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			c.handler.OnError(err)
			break
		}

		c.handler.OnReceive(message.Parse(string(line)))
	}
}

func (c *conn) sender() {
	for {
		select {
		case msg := <-c.in:
			c.handler.OnSend(msg)
			c.conn.Write([]byte(msg.String()))
		case <-c.quit:
			c.conn.Close()
			c.handler.OnQuit()
			break
		}
	}
}
