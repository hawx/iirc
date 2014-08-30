package server

import (
	"github.com/hawx/iirc/message"
	"github.com/hawx/iirc/errors"
	"log"
	"bufio"
	"net"
)

type Server struct {
	address string
	port    string
	name    string
	conn    net.Listener

	in       chan string
	quit     chan struct{}
	clients  *Clients
	channels *Channels
}

func NewServer(name, address, port string) *Server {
	in := make(chan string)
	quit := make(chan struct{})
	clients := NewClients()
	channels := NewChannels()

	return &Server{
		address:  address,
		port:     port,
		name:     name,
		in:       in,
  	quit:     quit,
		clients:  clients,
		channels: channels,
	}
}

func (s *Server) Name() string {
	return s.name
}

func (s *Server) Start() {
	tcp, err := net.Listen("tcp", s.address+":"+s.port)
	if err != nil {
		log.Fatal(err)
	}

	s.conn = tcp
	log.Println("server listening at", s.address+":"+s.port)
	defer tcp.Close()

	go s.receiver()

	for {
		conn, err := tcp.Accept()

		if err != nil {
			select {
			case <-s.quit:
				log.Println("stopping server")
				s.clients.Close()
				return
			default:
				log.Println(err)
			}

			continue
		}

		go s.accept(conn)
	}
}

func (s *Server) Stop() {
	close(s.quit)
	s.conn.Close()
}

func Start(name, address, port string) {
	s := NewServer(name, address, port)
	s.Start()
}

func (s *Server) Address() string {
	return s.address
}

func (s *Server) Names() []string {
	return s.clients.Names()
}

func (s *Server) Remove(c *Client) {
	s.clients.Remove(c)
}

func (s *Server) FindChannel(name string) *Channel {
	ch, _ := s.channels.Find(name)
	return ch
}

func (s *Server) Join(c *Client, channelName string) *Channel {
	ch, ok := s.channels.Find(channelName)

	if !ok {
		s.channels.Add(ch)
	}

	ch.Join(c)
	return ch
}

func (s *Server) Part(c *Client, channelName string) {
	ch, ok := s.channels.Find(channelName)

	if !ok {
		return
	}

	ch.Leave(c)

	if ch.Empty() {
		s.channels.Remove(ch)
	}
}

func (s *Server) List() []string {
	return s.channels.Names()
}

// accept handles the initial negotiation phase for connecting to the
// server. This is basically in 4 parts:
//
// * PASS
// * NICK
// * USER
//
// On success a RPL_WELCOME response is returned to the new client.
func (s *Server) accept(conn net.Conn) bool {
	r := bufio.NewReader(conn)

	line, err := r.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return false
	}

	parsed := message.Parse(string(line))

	if parsed.Command != "PASS" {
		return false
	}

	if !parsed.Params.Any() {
		conn.Write([]byte(errors.NeedMoreParams("PASS").String()))
		return false
	}

	password := parsed.Params.Get(0)
	if password != "test" {
		return false
	}

	client := NewClient("", conn, s)
	s.clients.Add(client)

	return true
}

func (s *Server) receiver() {
	// for {
	// 	select {
	// 	case msg := <-s.in:
	// 		// s.clients.Broadcast(msg)
	// 	}
	// }
}
