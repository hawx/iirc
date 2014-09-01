package test

import "testing"

func TestAway(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("AWAY :Be back in 10")
		assertResponse(t, a, "306 :You have been marked as being away\r\n")
	})
}

func TestAwayWithNoArguments(t *testing.T) {
	with(t, func(a *TestClient) {
		a.Authenticate()

		a.Send("AWAY")
		assertResponse(t, a, "305 :You are no longer marked as being away\r\n")
	})
}
