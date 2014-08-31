package test

import "testing"

func TestPing(t *testing.T) {
	with(t, func(client *TestClient) {
		client.Authenticate()

		client.Send("PING")
		assertResponse(t, client, "PONG "+address+"\r\n")

		client.Send("QUIT")
	})
}
