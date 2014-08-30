package message

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixWithServername(t *testing.T) {
	const serverName = "someserver"

	prefix := Prefix(serverName)
	if prefix.String() != serverName {
		t.Fail()
	}
}

func TestPrefixWithNickname(t *testing.T) {
	const nickname = "testnick"

	prefix := Prefix(nickname)
	if prefix.String() != nickname {
		t.Fail()
	}
}

func TestPrefixWithNicknameAndHost(t *testing.T) {
	const nickname = "testnick"
	const host = "testhost"

	prefix := Prefix(nickname, host)
	if prefix.String() != nickname + "@" + host {
		t.Fail()
	}
}

func TestPrefixWithNicknameAndUserAndHost(t *testing.T) {
	const nickname = "testnick"
	const user = "testuser"
	const host = "testhost"

	prefix := Prefix(nickname, user, host)
	if prefix.String() != nickname + "!" + user + "@" + host {
		t.Fail()
	}
}

func TestParsePrefixWithServername(t *testing.T) {
	const s = "iamserver"

	prefix := parsePrefix(s)
	assert.Equal(t, prefix.name, s)
	assert.Equal(t, prefix.user, "")
	assert.Equal(t, prefix.host, "")
}

func TestParsePrefixWithNicknameAndHost(t *testing.T) {
	const s = "nick@local"

	prefix := parsePrefix(s)
	assert.Equal(t, prefix.name, "nick")
	assert.Equal(t, prefix.user, "")
	assert.Equal(t, prefix.host, "local")
}

func TestParsePrefixWithNicknameAndUserAndHost(t *testing.T) {
	const s = "nick!what@local"

	prefix := parsePrefix(s)
	assert.Equal(t, prefix.name, "nick")
	assert.Equal(t, prefix.user, "what")
	assert.Equal(t, prefix.host, "local")
}
