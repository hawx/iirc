package test

import "testing"

func TestNick(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK test")

		a.Send("NICK changed")
		assertResponse(t, a, ":test NICK changed")
	})
}

func TestNickWithNoArgument(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS test")
		client.Send("NICK")
		assertResponse(t, client, ":test.irc 431 :No nickname given")
	})
}

func TestNickWithTakenNickname(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Send("PASS test")
		a.Send("NICK test")

		b.Send("PASS test")
		b.Send("NICK test")
		assertResponse(t, b, ":test.irc 433 test :Nickname is already in use")
	})
}
