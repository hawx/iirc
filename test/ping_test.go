package test

import "testing"

func TestPing(t *testing.T) {
	with(t, func(client *TestClient) {
		authenticate(client, "test", "user1", "Test User")

		client.Send("PING")
		assertResponse(t, client, "PONG "+address+"\r\n")

		client.Send("QUIT")
	})
}
