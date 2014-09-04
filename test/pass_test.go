package test

import "testing"

func TestPassWithNoArgument(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS")
		assertResponse(t, client, ":test.irc 461 PASS :Not enough parameters")
	})
}
