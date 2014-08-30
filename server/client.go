package server

import (
	"bufio"
	"github.com/hawx/iirc/message"
	"log"
	"net"
)

type Client struct {
	in       chan message.M
	out      chan string
	quit     chan struct{}
	Name     string
	userName string
	realName string
	mode     string
	conn     net.Conn
	server   *Server
	channels *Channels
	awayMsg  string
}

func NewClient(name string, conn net.Conn, s *Server) *Client {
	in := make(chan message.M)
	out := make(chan string)
	quit := make(chan struct{})

	client := &Client{
		in:       in,
		out:      out,
		quit:     quit,
		Name:     name,
		realName: "",
		conn:     conn,
		server:   s,
		channels: NewChannels(),
	}

	log.Println("client started")

	go client.receiver()
	go client.sender()

	return client
}

func (c *Client) Send(msg message.M) {
	c.in <- msg
}

func (c *Client) Close() {
	c.quit <- struct{}{}
}

func (c *Client) receiver() {
	r := bufio.NewReader(c.conn)

	log.Println("client receiving")

	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break
		}

		l := message.Parse(string(line))
		log.Println(c.Name, "->", l)

		switch l.Command {
		case "QUIT":
			c.Close()
			break

		case "PING":
			Ping(c, c.server)

		case "NICK":
			Nick(c, c.server, l.Args())

		case "USER":
			User(c, c.server, l.Args(), l.Command)

		case "NAMES":
			Names(c, c.server, l.Args())

		case "JOIN":
			Join(c, c.server, l.Args())

		case "PART":
			Part(c, c.server, l.Args())

		case "TOPIC":
			Topic(c, c.server, l.Args())
		}
	}
}

func (c *Client) sender() {
	for {
		select {
		case msg := <-c.in:
			log.Print(c.Name, "<-", msg.String())
			c.conn.Write([]byte(msg.String()))
		case <-c.quit:
			c.conn.Close()
			c.server.Remove(c)
			break
		}
	}
}
