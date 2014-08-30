package message

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestParseWithNoNewline(t *testing.T) {
	assert := assert.New(t)
	const line = "PING"

	message := Parse(line)

	assert.Nil(message.Prefix)
	assert.Equal(message.Command, "PING")
	assert.Empty(message.Params.All())
}

func TestParseCommand(t *testing.T) {
	assert := assert.New(t)
	const line = "PING\r\n"

	message := Parse(line)

	assert.Nil(message.Prefix)
	assert.Equal(message.Command, "PING")
	assert.Empty(message.Params.All())
}

func TestParseCommandWithSource(t *testing.T) {
	assert := assert.New(t)
	const line = ":heywhat PING\r\n"

	message := Parse(line)

	assert.NotNil(message.Prefix)
	assert.Equal(message.Prefix.name, "heywhat")
	assert.Equal(message.Prefix.host, "")
	assert.Equal(message.Prefix.user, "")
	assert.Equal(message.Command, "PING")
	assert.Empty(message.Params.l)
}

func TestParseCommandWithParams(t *testing.T) {
	assert := assert.New(t)
	const line = "JOIN #channel\r\n"

	message := Parse(line)

	assert.Nil(message.Prefix)
	assert.Equal(message.Command, "JOIN")
	assert.NotNil(message.Params)
	assert.Equal(len(message.Params.l), 1)
	assert.Equal(message.Params.l[0], "#channel")
	assert.Equal(message.Params.t, "")
}

func TestParseCommandWithFifteenParams(t *testing.T) {
	assert := assert.New(t)
	const line = "JOIN 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15\r\n"

	message := Parse(line)

	assert.Nil(message.Prefix)
	assert.Equal(message.Command, "JOIN")
	assert.NotNil(message.Params)

	assert.Equal(len(message.Params.l), 15)
	for i := 1; i <= 15; i++ {
		assert.Equal(message.Params.l[i-1], strconv.Itoa(i))
	}

	assert.Equal(message.Params.t, "")
}

func TestParseCommandWithParamsAndText(t *testing.T) {
	assert := assert.New(t)
	const line = "SOMETHING and :what you gonna do\r\n"

	message := Parse(line)

	assert.Nil(message.Prefix)
	assert.Equal(message.Command, "SOMETHING")
	assert.NotNil(message.Params)
	assert.Equal(len(message.Params.l), 1)
	assert.Equal(message.Params.l[0], "and")
	assert.Equal(message.Params.t, "what you gonna do")
}
