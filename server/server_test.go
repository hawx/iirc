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

	if !ok { t.Fatalf("expected: %#v, timed out", expect) }

	if resp != expect {
		t.Fatalf("expected: %#v, got: %#v", expect, resp)
	}
}

func expectResponse(t *testing.T, c *TestClient, expect string) {
	resp, ok := getResponse(c)

	if !ok { t.Errorf("expected: %#v, timed out", expect) }

	if resp != expect {
		t.Errorf("expected: %#v, got: %#v", expect, resp)
	}
}

func authenticate(c *TestClient, nickName, userName, realName string) {
	c.Send("PASS test")
	c.Send("NICK " + nickName)
	c.Send("USER " + userName + " 8 * :" + realName)
	getResponse(c)
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
}

func withClient(t *testing.T, f func(*TestClient)) {
	conn, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		t.Fatal(err)
	}

	client := NewTestClient(conn)
	f(client)
	client.Close()
}

func withServer(t *testing.T, f func(*Server)) {
	s := NewServer(serverName, address, port)
	go s.Start()
	time.Sleep(time.Millisecond)

	f(s)

	time.Sleep(time.Millisecond)
	s.Stop()
}

func with(t *testing.T, f func(*TestClient)) {
	withServer(t, func(s *Server) {
		withClient(t, func(c *TestClient) {
			f(c)
		})
	})
}

func with2(t *testing.T, f func(a, b *TestClient)) {
	withServer(t, func(s *Server) {
		withClient(t, func(a *TestClient) {
			withClient(t, func(b *TestClient) {
				f(a, b)
			})
		})
	})
}

// PASS

func TestPassWithNoArgument(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS")
		assertResponse(t, client, "461 PASS :Not enough parameters\r\n")
	})
}

// NICK

func TestNick(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK test")

		a.Send("NICK changed")
		assertResponse(t, a, ":test NICK changed\r\n")
	})
}

func TestNickWithNoArgument(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS test")
		client.Send("NICK")
		assertResponse(t, client, "431 :No nickname given\r\n")
	})
}

func TestNickWithTakenNickname(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Send("PASS test")
		a.Send("NICK test")

		b.Send("PASS test")
		b.Send("NICK test")
		assertResponse(t, b, "433 test :Nickname is already in use\r\n")
	})
}

// USER

func TestUser(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER testuser 0 * :Mr Test")

		assertResponse(t, a, ":"+serverName+" 001 testuser\r\n")
	})
}

func TestUserWithNoArgument(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER")
		assertResponse(t, a, "461 USER :Not enough parameters\r\n")
	})
}

func TestUserWhenAlreadySent(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER testuser 0 * :Mr Test")
		getResponse(a)

		a.Send("USER testuser 0 * :Mr Test")
		assertResponse(t, a, "462 :Unauthorized command (already registered)\r\n")
	})
}

// PING

func TestPing(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS test")
		client.Send("NICK josh")
		client.Send("USER test test")
		getResponse(client)

		client.Send("PING")
		assertResponse(t, client, "PONG "+address+"\r\n")

		client.Send("QUIT")
	})
}

// NAMES

func TestNames(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS test")
		client.Send("NICK josh")
		client.Send("USER user1 8 * :Test User")
		getResponse(client)

		client.Send("NAMES #test")
		assertResponse(t, client, ":test.irc 353 josh = #test\r\n")
		assertResponse(t, client, ":test.irc 366 josh #test :End of /NAMES list.\r\n")

		client.Send("QUIT")
	})
}

func TestNamesWithPersonInChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Send("PASS test")
		a.Send("NICK a")
		a.Send("USER user1 8 * :Test User")
		getResponse(a)

		b.Send("PASS test")
		b.Send("NICK b")
		b.Send("USER user2 8 * :Test User2")
		getResponse(b)

		a.Send("JOIN #test")

		b.Send("NAMES #test")
		assertResponse(t, b, ":test.irc 353 b = #test :a\r\n")
		assertResponse(t, b, ":test.irc 366 b #test :End of /NAMES list.\r\n")

		a.Send("QUIT")
		b.Send("QUIT")
	})
}

func TestNamesWithPersonAndSelfInChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Send("PASS test")
		a.Send("NICK a")
		a.Send("USER user1 8 * :Test User")
		getResponse(a)

		b.Send("PASS test")
		b.Send("NICK b")
		b.Send("USER user2 8 * :Test User2")
		getResponse(b)

		a.Send("JOIN #test")
		getResponse(a)
		b.Send("JOIN #test")
		getResponse(b)

		b.Send("NAMES #test")
		assertResponse(t, b, ":test.irc 353 b = #test :a b\r\n")
		assertResponse(t, b, ":test.irc 366 b #test :End of /NAMES list.\r\n")

		a.Send("QUIT")
		b.Send("QUIT")
	})
}


// JOIN

func TestJoin(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK a")
		a.Send("USER user1 8 * :Test User")
		getResponse(a)

		a.Send("JOIN #test")
		assertResponse(t, a, ":a!user1@test.irc JOIN :#test\r\n")
	})
}

func TestJoinIsBroadcastToChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Send("PASS test")
		a.Send("NICK a")
		a.Send("USER user1 8 * :Test User")
		getResponse(a)

		b.Send("PASS test")
		b.Send("NICK b")
		b.Send("USER user2 8 * :Test User 2")
		getResponse(b)

		a.Send("JOIN #test")
		assertResponse(t, a, ":a!user1@test.irc JOIN :#test\r\n")

		b.Send("JOIN #test")
		assertResponse(t, a, ":b!user2@test.irc JOIN :#test\r\n")
		assertResponse(t, b, ":b!user2@test.irc JOIN :#test\r\n")
	})
}

// TOPIC

func TestTopic(t *testing.T) {
	with(t, func(a *TestClient) {
		authenticate(a, "test", "user1", "Test User")

		a.Send("TOPIC #test")
		assertResponse(t, a, "331 #test :No topic is set\r\n")

		a.Send("QUIT")
	})
}

func TestSetTopic(t *testing.T) {
	with(t, func(a *TestClient) {
		authenticate(a, "test", "user1", "Test User")

		a.Send("JOIN #test")
		getResponse(a)

		a.Send("TOPIC #test :Cool stufff only")
		assertResponse(t, a, ":test!user1@test.irc TOPIC #test :Cool stufff only\r\n")

		a.Send("TOPIC #test")
		assertResponse(t, a, "332 #test :Cool stufff only\r\n")

		a.Send("QUIT")
	})
}

// PART

func TestPart(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		authenticate(a, "a", "user1", "Test User")
		authenticate(b, "b", "user2", "Test User2")

		a.Send("JOIN #test")
		getResponse(a)

		b.Send("JOIN #test")
		getResponse(a)
		getResponse(b)

		a.Send("PART #test")
		assertResponse(t, a, ":a!user1@test.irc PART #test :a\r\n")
		assertResponse(t, b, ":a!user1@test.irc PART #test :a\r\n")

		a.Send("QUIT")
		b.Send("QUIT")
	})
}

func TestPartWhenNoSuchChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		authenticate(a, "a", "user1", "Test User")

		a.Send("PART #test")
		assertResponse(t, a, "442 #test :You're not on that channel\r\n")

		a.Send("QUIT")
	})
}
