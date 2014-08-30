package test

import "testing"

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
