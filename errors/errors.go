package errors

import "github.com/hawx/iirc/message"

const (
	// Returned when a nickname parameter expected for a command and isn't found.
	//
	// ":No nickname given"
	ERR_NONICKNAMEGIVEN = "431"

	// Returned when a NICK message is processed that results in an attempt to
	// change to a currently existing nickname.
	//
	// "<nick> :Nickname is already in use"
  ERR_NICKNAMEINUSE = "433"

	// Returned by the server whenever a client tries to perform a channel
	// affecting command for which the client isn't a member.
	//
	// "<channel> :You're not on that channel"
	ERR_NOTONCHANNEL = "442"

	// Returned by the server by numerous commands to indicate to the client that
	// it didn't supply enough parameters.
	//
	// "<command> :Not enough parameters"
	ERR_NEEDMOREPARAMS = "461"

	// Returned by the server to any link which tries to change part of the
	// registered details (such as password or user details from second USER
	// message).
	//
	// ":Unauthorized command (already registered)"
	ERR_ALREADYREGISTRED = "462"
)

func NoNicknameGiven() message.M {
	return message.MessageParams(
		ERR_NONICKNAMEGIVEN,
		message.ParamsT([]string{}, "No nickname given"))
}

func NicknameInUse(nick string) message.M {
	return message.MessageParams(
		ERR_NICKNAMEINUSE,
		message.ParamsT([]string{nick}, "Nickname is already in use"))
}

func NotOnChannel(channel string) message.M {
	return message.MessageParams(
		ERR_NOTONCHANNEL,
		message.ParamsT([]string{channel}, "You're not on that channel"))
}

func NeedMoreParams(command string) message.M {
	return message.MessageParams(
		ERR_NEEDMOREPARAMS,
		message.ParamsT([]string{command}, "Not enough parameters"))
}

func AlreadyRegistered() message.M {
	return message.MessageParams(
		ERR_ALREADYREGISTRED,
		message.ParamsT([]string{}, "Unauthorized command (already registered)"))
}
