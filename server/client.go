package server

import (
	"bufio"
	"github.com/hawx/iirc/channel"
	"github.com/hawx/iirc/handler"
	"github.com/hawx/iirc/message"
	"log"
	"net"
)

type Client struct {
	in       chan message.M
	out      chan string
	quit     chan struct{}
	name     string
	userName string
	realName string
	mode     string
	conn     net.Conn
	server   *Server
	channels *channel.Channels
	awayMsg  string
}

func (c *Client) Name() string     { return c.name }
func (c *Client) SetName(n string) { c.name = n }

func (c *Client) UserName() string     { return c.userName }
func (c *Client) SetUserName(n string) { c.userName = n }

func (c *Client) SetMode(n string) { c.mode = n }

func (c *Client) RealName() string     { return c.realName }
func (c *Client) SetRealName(n string) { c.realName = n }

func (c *Client) Channels() *channel.Channels { return c.channels }

func NewClient(name string, conn net.Conn, s *Server) *Client {
	in := make(chan message.M)
	out := make(chan string)
	quit := make(chan struct{})

	client := &Client{
		in:       in,
		out:      out,
		quit:     quit,
		name:     name,
		realName: "",
		conn:     conn,
		server:   s,
		channels: channel.NewChannels(),
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
			c.Send(message.MessageParams(
				"ERROR",
				message.ParamsT([]string{}, "Closing Link: "+c.Name())))

			c.Channels().Each(func(ch *channel.Channel) {
				ch.Broadcast(message.MessagePrefix(
					message.Prefix(c.Name(), c.UserName(), c.server.Name()),
					"QUIT"))
			})

			c.Close()
			break

		case "PING":
			handler.Ping(c, c.server)

		case "NICK":
			handler.Nick(c, c.server, l.Args())

		case "USER":
			handler.User(c, c.server, l.Args(), l.Command)

		case "NAMES":
			handler.Names(c, c.server, l.Args())

		case "JOIN":
			handler.Join(c, c.server, l.Args())

		case "PART":
			handler.Part(c, c.server, l.Args())

		case "TOPIC":
			handler.Topic(c, c.server, l.Args())
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
