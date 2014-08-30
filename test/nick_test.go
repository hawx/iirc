package test

import "testing"

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
