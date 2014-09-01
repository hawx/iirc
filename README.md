# iirc

> Do not put this on a server, it is not ready.

Implements a (subset of irc)-server in the style of
<https://github.com/sstephenson/hector>. Like __hector__ the supported commands
are:

- [X] USER and PASS -- Authenticates you to the server. (Your client sends these as
  soon as it connects.)
- [X] NICK -- Sets your nickname.
- [X] JOIN -- Joins a channel.
- [X] PRIVMSG and NOTICE -- Sends a message to another nickname or channel.
- [X] TOPIC -- Changes or returns the topic of a channel.
- [X] NAMES -- Shows a list of which nicknames are on a channel.
- [ ] WHO -- Like NAMES, but returns more information. (Your client probably sends
  this when it joins a channel.)
- [ ] WHOIS -- Shows information about a nickname, including how long it has been
  connected.
- [X] PART -- Leaves a channel.
- [X] AWAY -- Marks or unmarks you as being away.
- [ ] INVITE -- Invites another user to a channel.
- [X] PING -- (Your client uses this command to measure the speed of its connection
  to the server.)
- [X] QUIT -- Disconnects from the server.
