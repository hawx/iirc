package server

import (
	"log"
	"net"

	"hawx.me/code/iirc/channel"
	"hawx.me/code/iirc/commands"
	"hawx.me/code/iirc/connection"
	"hawx.me/code/iirc/message"
)

type Client struct {
	conn     connection.Conn
	name     string
	userName string
	realName string
	mode     string
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

func (c *Client) AwayMessage() string     { return c.awayMsg }
func (c *Client) SetAwayMessage(n string) { c.awayMsg = n }

func NewClient(name string, conn net.Conn, s *Server) *Client {
	client := &Client{
		name:     name,
		realName: "",
		server:   s,
		channels: channel.NewChannels(),
	}

	handler := connection.Log(name, clientHandler{client})
	client.conn = connection.NewConn(conn, handler)

	log.Println("client started")

	return client
}

func (c *Client) Send(msg message.M) {
	c.conn.Send(msg)
}

func (c *Client) SendExcept(msg message.M, name string) {
	if c.Name() != name {
		c.conn.Send(msg)
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

type clientHandler struct {
	client *Client
}

var handlers = map[string]commands.Command{
	"PING":    commands.Ping,
	"NICK":    commands.Nick,
	"USER":    commands.User,
	"NAMES":   commands.Names,
	"JOIN":    commands.Join,
	"PART":    commands.Part,
	"TOPIC":   commands.Topic,
	"PRIVMSG": commands.PrivMsg,
	"NOTICE":  commands.Notice,
	"AWAY":    commands.Away,
	"MODE":    commands.Mode,
	"WHO":     commands.Who,
	"WHOIS":   commands.Whois,
	"INVITE":  commands.Invite,
}

func (c clientHandler) OnReceive(l message.M) {
	if l.Command == "QUIT" {
		c.client.Send(message.Message3(
			nil,
			"ERROR",
			message.ParamsT([]string{}, "Closing Link: "+c.client.Name())))

		c.client.Channels().Each(func(ch *channel.Channel) {
			ch.Send(message.Message3(
				message.Prefix(c.client.Name(), c.client.UserName(), c.client.server.Name()),
				"QUIT",
				nil))
		})

		c.client.Close()
		return
	}

	if handler, ok := handlers[l.Command]; ok {
		handler(c.client, c.client.server, l.Args())
	}
}

func (c clientHandler) OnSend(m message.M) {}
func (c clientHandler) OnError(e error)    {}

func (c clientHandler) OnQuit() {
	c.client.server.Remove(c.client)
}
