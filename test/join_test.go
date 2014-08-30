package test

import "testing"

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
