package test

import "testing"

func TestQuit(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		authenticate(a, "a", "user1", "Test User 1")
		authenticate(b, "b", "user2", "Test User 2")

		a.Send("JOIN #test")
		getResponse(a)
		a.Send("JOIN #other")
		getResponse(a)

		b.Send("JOIN #test")
		getResponse(a)
		getResponse(b)
		b.Send("JOIN #other")
		getResponse(a)
		getResponse(b)

		a.Send("QUIT")
		assertResponse(t, b, ":a!user1@test.irc QUIT\r\n")
		assertResponse(t, a, "ERROR :Closing Link: a\r\n")
	})
}
