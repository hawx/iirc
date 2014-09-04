package test

import "testing"

func TestMode(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("MODE " + a.nickName)
		assertResponse(t, a, ":test.irc 221 " + a.nickName + " +")
	})
}

func TestModeWithNoArguemtn(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("MODE")
		assertResponse(t, a, ":test.irc 461 MODE :Not enough parameters")
	})
}

func TestModeWithChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("MODE #test")
		assertResponse(t, a, ":test.irc 324 " + a.nickName + " #test +")
	})
}

func TestModeWithChannelRetrieveBans(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		a.Send("MODE #test +b")
		assertResponse(t, a, ":test.irc 368 " + a.nickName + " #test :End of Channel Ban List")
	})
}
