package test

import "testing"

func TestPrivMsgToUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		const msg = "Hey I have a message"
		a.Send("PRIVMSG " + b.nickName + " :" + msg)
		assertResponse(t, b, a.Prefix() + " PRIVMSG " + b.nickName + " :" + msg + "\r\n")
	})
}

func TestPrivMsgToAwayUser(t *testing.T) {
	with2(t, func(a, b *TestClient) {
		a.Authenticate()
		b.Authenticate()

		const awayMsg = "Be back in 5"
		b.Send("AWAY :" + awayMsg)
		getResponse(b)

		const msg = "Hey I have a message"
		a.Send("PRIVMSG " + b.nickName + " :" + msg)
		assertResponse(t, b, a.Prefix() + " PRIVMSG " + b.nickName + " :" + msg + "\r\n")
		assertResponse(t, a, "301 " + b.nickName + " :" + awayMsg + "\r\n")
	})
}

func TestPrivMsgToChannel(t *testing.T) {
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

		msg := "Hey I have a message"
		a.Send("PRIVMSG #test :" + msg)
		assertResponse(t, b, a.Prefix() + " PRIVMSG #test :" + msg + "\r\n")

		msg = "Ok"
		b.Send("PRIVMSG #test :" + msg)
		assertResponse(t, a, b.Prefix() + " PRIVMSG #test :" + msg + "\r\n")
	})
}
