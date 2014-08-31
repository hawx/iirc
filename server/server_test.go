package server

import (
	"bufio"
	"net"
	"testing"
	"time"
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
}

func NewTestClient(conn net.Conn) *TestClient {
	client := &TestClient{
		in:   make(chan string),
		Out:  make(chan string),
		err:  make(chan error),
		quit: make(chan struct{}),
		conn: conn,
	}

	go client.receiver()
	go client.sender()

	return client
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

func assertResponse(t *testing.T, c *TestClient, expect string) {
	resp, ok := getResponse(c)

	if !ok {
		t.Fatalf("expected: %#v, timed out", expect)
	}

	if resp != expect {
		t.Fatalf("expected: %#v, got: %#v", expect, resp)
	}
}

func TestServer(t *testing.T) {
	s := NewServer(serverName, address, port)
	go s.Start()
	time.Sleep(time.Millisecond)

	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(conn)

	client.Send("PASS test")
	client.Send("NICK josh")
	client.Send("USER user1 0 * :Test User")
	assertResponse(t, client, ":"+serverName+" 001 josh\r\n")

	client.Send("QUIT")
	client.Close()
	s.Stop()
	time.Sleep(time.Millisecond)
}
