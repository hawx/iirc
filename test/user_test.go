package test

import "testing"

func TestUser(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER testuser 0 * :Mr Test")

		assertResponse(t, a, ":"+serverName, "001 testuser")
	})
}

func TestUserWithNoArgument(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER")
		assertResponse(t, a, ":test.irc 461 USER :Not enough parameters")
	})
}

func TestUserWhenAlreadySent(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Send("PASS test")
		a.Send("NICK testuser")
		a.Send("USER testuser 0 * :Mr Test")
		getResponse(a)

		a.Send("USER testuser 0 * :Mr Test")
		assertResponse(t, a, ":test.irc 462 :Unauthorized command (already registered)")
	})
}
