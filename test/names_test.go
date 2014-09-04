package test

import "testing"

func TestNames(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("NAMES #test")
		assertResponse(t, a, ":test.irc 353 "+a.nickName+" = #test")
		assertResponse(t, a, ":test.irc 366 "+a.nickName+" #test :End of /NAMES list.")

		a.Send("QUIT")
	})
}

func TestNamesWithPersonInChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)

		b.Send("NAMES #test")
		assertResponse(t, b, ":test.irc 353 "+b.nickName+" = #test :"+a.nickName)
		assertResponse(t, b, ":test.irc 366 "+b.nickName+" #test :End of /NAMES list.")
	})
}

func TestNamesWithPersonAndSelfInChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)
		b.Send("JOIN #test")
		getResponse(b)
		getResponse(b)

		b.Send("NAMES #test")
		assertResponse(t, b, ":test.irc 353 "+b.nickName+" = #test :"+a.nickName+" "+b.nickName)
		assertResponse(t, b, ":test.irc 366 "+b.nickName+" #test :End of /NAMES list.")
	})
}
