package test

import "testing"

func TestWhois(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("WHOIS missing")
		assertResponse(t, a, ":test.irc 401 missing :No such nick/channel")
		assertResponse(t, a, ":test.irc 318 missing :End of WHOIS list")
	})
}

func TestWhoisWithUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("WHOIS " + b.nickName)
		assertResponse(t, a, ":test.irc 311", b.nickName, b.userName, "test.irc *", ":"+b.realName)
		assertResponse(t, a, ":test.irc 312", b.nickName, "test.irc")
		assertResponse(t, a, ":test.irc 317", b.nickName, "0 :seconds idle")
		assertResponse(t, a, ":test.irc 318", b.nickName, ":End of WHOIS list")
	})
}

func TestWhoisWithUserOnChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		b.Send("JOIN #test")

		a.Send("WHOIS " + b.nickName)
		assertResponse(t, a, ":test.irc 311", b.nickName, b.userName, "test.irc *", ":"+b.realName)
		assertResponse(t, a, ":test.irc 312", b.nickName, "test.irc")
		assertResponse(t, a, ":test.irc 319", b.nickName, ":#test")
		assertResponse(t, a, ":test.irc 317", b.nickName, "0 :seconds idle")
		assertResponse(t, a, ":test.irc 318", b.nickName, ":End of WHOIS list")
	})
}

func TestWhoisWithAwayUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		b.Send("AWAY :Gone out")

		a.Send("WHOIS " + b.nickName)
		assertResponse(t, a, ":test.irc 301", b.nickName, ":Gone out")
		assertResponse(t, a, ":test.irc 311", b.nickName, b.userName, "test.irc *", ":"+b.realName)
		assertResponse(t, a, ":test.irc 312", b.nickName, "test.irc")
		assertResponse(t, a, ":test.irc 317", b.nickName, "0 :seconds idle")
		assertResponse(t, a, ":test.irc 318", b.nickName, ":End of WHOIS list")
	})
}
