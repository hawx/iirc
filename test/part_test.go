package test

import "testing"

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
