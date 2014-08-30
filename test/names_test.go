package test

import "testing"

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
