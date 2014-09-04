package test

import "testing"

func TestWhoWithMissingParam(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("WHO")
		assertResponse(t, a, ":test.irc 461 WHO :Not enough parameters")
	})
}

func TestWhoWithUndefinedChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("WHO #test")
		assertResponse(t, a, ":test.irc 315 " + a.nickName + " #test")
	})
}

func TestWhoWithActiveChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")

		b.Send("WHO #test")
		assertResponse(t, b, ":test.irc 352 " + b.nickName + " #test " + a.userName + " test.irc test.irc " + a.nickName + " H :0 " + a.realName)
		assertResponse(t, b, ":test.irc 315 " + b.nickName + " #test")
	})
}

func TestWhoWithJoinedChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)
		b.Send("JOIN #test")
		getResponse(a)

		a.Send("WHO #test")
		assertResponse(t, a, ":test.irc 352 " + a.nickName + " #test " + a.userName + " test.irc test.irc " + a.nickName + " H :0 " + a.realName)
		assertResponse(t, a, ":test.irc 352 " + a.nickName + " #test " + b.userName + " test.irc test.irc " + b.nickName + " H :0 " + b.realName)
		assertResponse(t, a, ":test.irc 315 " + a.nickName + " #test")
	})
}

func TestWhoForUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("WHO " + b.nickName)
		assertResponse(t, a, ":test.irc 352 " + a.nickName + " " + b.nickName + " " + b.userName + " test.irc test.irc " + b.nickName + " H :0 " + b.realName)
		assertResponse(t, a, ":test.irc 315 " + a.nickName + " " + b.nickName)
	})
}

func TestWhoForUndefinedUser(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("WHO what")
		assertResponse(t, a, ":test.irc 315 " + a.nickName + " what")
	})
}
