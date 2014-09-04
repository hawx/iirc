package test

import "testing"

func TestTopic(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("TOPIC #test")
		assertResponse(t, a, ":test.irc 331 #test :No topic is set")
	})
}

func TestTopicWithNoParams(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("TOPIC")
		assertResponse(t, a, ":test.irc 461 TOPIC :Not enough parameters")
	})
}

func TestSetTopic(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("TOPIC #test :Cool stufff only")
		assertResponse(t, a, a.Prefix(), "TOPIC #test :Cool stufff only")

		a.Send("TOPIC #test")
		assertResponse(t, a, ":test.irc 332 #test :Cool stufff only")
	})
}
