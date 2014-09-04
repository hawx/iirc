package test

import "testing"

func TestJoin(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix(), "JOIN :#test")
		assertResponse(t, a, ":test.irc 331 #test :No topic is set")
	})
}

func TestJoinWithNoArguments(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN")
		assertResponse(t, a, ":test.irc 461 JOIN :Not enough parameters")
	})
}

func TestJoinIsBroadcastToChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix()+" JOIN :#test")
		assertResponse(t, a, ":test.irc 331 #test :No topic is set")

		b.Send("JOIN #test")
		assertResponse(t, a, b.Prefix()+" JOIN :#test")
		assertResponse(t, b, b.Prefix()+" JOIN :#test")
		assertResponse(t, b, ":test.irc 331 #test :No topic is set")
	})
}

func TestJoinWhenJoiningMultiple(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test,#what")
		assertResponse(t, a, a.Prefix()+" JOIN :#test")
		assertResponse(t, a, ":test.irc 331 #test :No topic is set")
		assertResponse(t, a, a.Prefix()+" JOIN :#what")
		assertResponse(t, a, ":test.irc 331 #what :No topic is set")
	})
}
