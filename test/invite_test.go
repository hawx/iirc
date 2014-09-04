package test

import "testing"

func TestInvite(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("INVITE " + b.nickName + " #test")
		assertResponse(t, b, a.Prefix(), "INVITE", b.nickName, ":#test")
	})
}

func TestInvitieWithNoChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("INVITE missing")
		assertResponse(t, a, ":test.irc 461 INVITE :Not enough parameters")
	})
}

func TestInvitieWithNoUserOrChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("INVITE")
		assertResponse(t, a, ":test.irc 461 INVITE :Not enough parameters")
	})
}

func TestInviteWithUndefinedUser(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("INVITE missing #test")
		assertResponse(t, a, ":test.irc 401 missing :No such nick/channel")
	})
}

func TestInviteWhenUserAlreadyOnChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		b.Send("JOIN #test")
		getResponse(a)

		a.Send("INVITE " + b.nickName + " #test")
		assertResponse(t, a, ":test.irc 443", b.nickName, "#test :is already on channel")
	})
}

func TestInviteWhenNotOnChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("INVITE " + b.nickName + " #test")
		assertResponse(t, a, ":test.irc 442 #test :You're not on that channel")
	})
}
