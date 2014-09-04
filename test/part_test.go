package test

import "testing"

func TestPart(t *testing.T) {
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

		a.Send("PART #test")
		assertResponse(t, a, a.Prefix()+" PART #test :"+a.nickName)
		assertResponse(t, b, a.Prefix()+" PART #test :"+a.nickName)
	})
}

func TestPartWhenNoSuchChannel(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("PART #test")
		assertResponse(t, a, ":test.irc 442 #test :You're not on that channel")
	})
}

func TestPartWhenNoParameters(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("PART")
		assertResponse(t, a, ":test.irc 461 PART :Not enough parameters")
	})
}
