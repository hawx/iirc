package message

import (
	"testing"
	"strconv"
	"strings"
)

func TestMessageWithOnlyCommand(t *testing.T) {
	message := Message("PING")

	if message.String() != "PING\r\n" {
		t.Error(message.String())
	}
}

func TestMessageWithPrefixAndCommand(t *testing.T) {
	message := MessagePrefix(Prefix("a.irc"), "PING")

	if message.String() != ":a.irc PING\r\n" {
		t.Error(message.String())
	}
}

func TestMessageWithCommandAndParams(t *testing.T) {
	message := MessageParams("PING", Params([]string{"who"}))

	if message.String() != "PING who\r\n" {
		t.Error(message.String())
	}
}

func TestMessageWithPrefixAndCommandAndParams(t *testing.T) {
	message := Message3(Prefix("testuser"), "NICK", Params([]string{"other"}))

	if message.String() != ":testuser NICK other\r\n" {
		t.Error(message.String())
	}
}

func TestMessageWithOver512Characters(t *testing.T) {
	ps := make([]string, 71)
	for i := 0; i < 71; i++ {
		ps[i] = "user" + strconv.Itoa(i + 1)
	}

	message := Message3(Prefix("test.irc"), "353 user = #test", ParamsT([]string{}, strings.Join(ps, " ")))

	parts := message.Parts()

	if len(parts) != 2 {
		t.Error("expected 2 parts")
	}

	expected := ":test.irc 353 user = #test :user1 user2 user3 user4 user5 user6 user7 user8 user9 user10 user11 user12 user13 user14 user15 user16 user17 user18 user19 user20 user21 user22 user23 user24 user25 user26 user27 user28 user29 user30 user31 user32 user33 user34 user35 user36 user37 user38 user39 user40 user41 user42 user43 user44 user45 user46 user47 user48 user49 user50 user51 user52 user53 user54 user55 user56 user57 user58 user59 user60 user61 user62 user63 user64 user65 user66 user67 user68 user69 user70\r\n"
	if parts[0].String() != expected {
		t.Errorf("\n  expected: %#v\n       got: %#v\n", expected, parts[0].String())
	}

	expected = ":test.irc 353 user = #test :user71\r\n"
	if parts[1].String() != expected {
		t.Errorf("\n  expected: %#v\n       got: %#v\n", expected, parts[1].String())
	}
}
