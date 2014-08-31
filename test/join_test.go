package test

import "testing"

func TestJoin(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, a, "331 #test :No topic is set\r\n")
	})
}

func TestJoinWithNoArguments(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN")
		assertResponse(t, a, "461 JOIN :Not enough parameters\r\n")
	})
}

func TestJoinIsBroadcastToChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, a, "331 #test :No topic is set\r\n")

		b.Send("JOIN #test")
		assertResponse(t, a, b.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, b, b.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, b, "331 #test :No topic is set\r\n")
	})
}

func TestJoinWhenJoiningMultiple(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test,#what")
		assertResponse(t, a, a.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, a, "331 #test :No topic is set\r\n")
		assertResponse(t, a, a.Prefix()+" JOIN :#what\r\n")
		assertResponse(t, a, "331 #what :No topic is set\r\n")
	})
}
