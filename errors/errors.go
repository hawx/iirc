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

	// Returned when a client tries to invite a user to a
	// channel they are already on.
	//
	// "<user> <channel> :is already on channel"
  ERR_USERONCHANNEL = "443"

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


	// Used to indicate the nickname parameter supplied to a command is currently
	// unused.
	//
	// "<nickname> :No such nick/channel"
	ERR_NOSUCHNICK = "401"
)

func err(host, code string, params []string, text string) message.M {
	return message.Message3(
		message.Prefix(host),
		code,
		message.ParamsT(params, text))
}

var empty = []string{}

func NoNicknameGiven(host string) message.M {
	return err(host, ERR_NONICKNAMEGIVEN, empty, "No nickname given")
}

func NicknameInUse(host, nick string) message.M {
	return err(host, ERR_NICKNAMEINUSE, []string{nick}, "Nickname is already in use")
}

func NotOnChannel(host, channel string) message.M {
	return err(host, ERR_NOTONCHANNEL, []string{channel}, "You're not on that channel")
}

func NeedMoreParams(host, command string) message.M {
	return err(host, ERR_NEEDMOREPARAMS, []string{command}, "Not enough parameters")
}

func AlreadyRegistered(host string) message.M {
	return err(host, ERR_ALREADYREGISTRED, empty, "Unauthorized command (already registered)")
}

func NoSuchNick(host, nick string) message.M {
	return err(host, ERR_NOSUCHNICK, []string{nick}, "No such nick/channel")
}

func UserOnChannel(host, nick, channel string) message.M {
	return err(host, ERR_USERONCHANNEL, []string{nick, channel}, "is already on channel")
}
