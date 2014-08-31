package test

import "testing"

func TestJoin(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix()+" JOIN :#test\r\n")
	})
}

func TestJoinIsBroadcastToChannel(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		a.Send("JOIN #test")
		assertResponse(t, a, a.Prefix()+" JOIN :#test\r\n")

		b.Send("JOIN #test")
		assertResponse(t, a, b.Prefix()+" JOIN :#test\r\n")
		assertResponse(t, b, b.Prefix()+" JOIN :#test\r\n")
	})
}
