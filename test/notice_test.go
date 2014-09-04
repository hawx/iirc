package test

import "testing"

func TestNoticeWithNoParams(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("NOTICE")
		assertResponse(t, a, ":test.irc 461 NOTICE :Not enough parameters")
	})
}

func TestNoticeWithNoMessage(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("NOTICE missing")
		assertResponse(t, a, ":test.irc 461 NOTICE :Not enough parameters")
	})
}

func TestNoticeToUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		const msg = "Hey I have a message"
		a.Send("NOTICE " + b.nickName + " :" + msg)
		assertResponse(t, b, a.Prefix() + " NOTICE " + b.nickName + " :" + msg)
	})
}

func TestNoticeToChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)
		b.Send("JOIN #test")
		getResponse(a)
		getResponse(b)
		getResponse(b)

		msg := "Hey I have a message"
		a.Send("NOTICE #test :" + msg)
		assertResponse(t, b, a.Prefix() + " NOTICE #test :" + msg)

		msg = "Ok"
		b.Send("NOTICE #test :" + msg)
		assertResponse(t, a, b.Prefix() + " NOTICE #test :" + msg)
	})
}
