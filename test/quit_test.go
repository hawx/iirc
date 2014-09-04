package test

import "testing"

func TestQuit(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		getResponse(a)
		getResponse(a)
		a.Send("JOIN #other")
		getResponse(a)
		getResponse(a)

		b.Send("JOIN #test")
		getResponse(a)
		getResponse(b)
		getResponse(b)
		b.Send("JOIN #other")
		getResponse(a)
		getResponse(b)
		getResponse(b)

		a.Send("QUIT")
		assertResponse(t, b, a.Prefix()+" QUIT")
		assertResponse(t, a, "ERROR :Closing Link: "+a.nickName)
	})
}
