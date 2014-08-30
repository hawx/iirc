package test

import "testing"

func TestPassWithNoArgument(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Send("PASS")
		assertResponse(t, client, "461 PASS :Not enough parameters\r\n")
	})
}
