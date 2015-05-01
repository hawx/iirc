package test

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"hawx.me/code/iirc/server"
)

const (
	address    = "127.0.0.1"
	port       = "6767"
	serverName = "test.irc"
)

type TestClient struct {
	in   chan string
	Out  chan string
	err  chan error
	quit chan struct{}
	conn net.Conn

	nickName string
	userName string
	realName string
}

var i int = 0

func NewTestClient(conn net.Conn) *TestClient {
	s := strconv.Itoa(i)
	i += 1

	client := &TestClient{
		in:       make(chan string),
		Out:      make(chan string),
		err:      make(chan error),
		quit:     make(chan struct{}),
		conn:     conn,
		nickName: "nick" + s,
		userName: "user" + s,
		realName: "Test User " + s,
	}

	go client.receiver()
	go client.sender()

	return client
}

func (c *TestClient) Prefix() string {
	return ":" + c.nickName + "!" + c.userName + "@" + serverName
}

func (c *TestClient) Authenticate() {
	c.Send("PASS test")
	c.Send("NICK " + c.nickName)
	c.Send("USER " + c.userName + " 8 * :" + c.realName)
	getResponse(c)
}

func (c *TestClient) Close() {
	c.quit <- struct{}{}
}

func (c *TestClient) Send(msg string) {
	c.in <- msg
}

func (c *TestClient) receiver() {
	r := bufio.NewReader(c.conn)

	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			c.err <- err
			break
		}

		c.Out <- string(line)
	}
}

func (c *TestClient) sender() {
	for {
		select {
		case buf := <-c.in:
			c.conn.Write([]byte(buf + "\r\n"))
			time.Sleep(time.Millisecond)
		case <-c.quit:
			c.conn.Close()
			break
		}
	}
}

func getResponse(c *TestClient) (string, bool) {
	select {
	case resp := <-c.Out:
		return resp, true
	case <-time.After(time.Second):
		return "", false
	}
}

func assertResponse(t *testing.T, c *TestClient, expects ...string) {
	resp, ok := getResponse(c)
	expect := strings.Join(expects, " ") + "\r\n"

	if !ok {
		t.Fatalf("expected: %#v, timed out", expect)
	}

	if resp != expect {
		t.Fatalf("expected: %#v, got: %#v", expect, resp)
	}
}

func withClient(t *testing.T, f func(*TestClient)) {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(conn)
	f(client)
	client.Send("QUIT")
	client.Close()
}

func withServer(t *testing.T, f func(*server.Server)) {
	s := server.NewServer(serverName, address, port)
	go s.Start()
	time.Sleep(time.Millisecond * 3)
	f(s)

	time.Sleep(time.Millisecond * 3)
	s.Stop()
}

func with(t *testing.T, f func(*TestClient)) {
	withServer(t, func(s *server.Server) {
		withClient(t, func(c *TestClient) {
			f(c)
		})
	})
}

func with2(t *testing.T, f func(a, b *TestClient)) {
	withServer(t, func(s *server.Server) {
		withClient(t, func(a *TestClient) {
			withClient(t, func(b *TestClient) {
				f(a, b)
			})
		})
	})
}
