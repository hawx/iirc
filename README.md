# iirc

Implements a (subset of irc)-server in the style of
<https://github.com/sstephenson/hector>. Like __hector__ the supported commands
are:

- [ ] USER and PASS -- Authenticates you to the server. (Your client sends these as
  soon as it connects.)
- [ ] NICK -- Sets your nickname.
- [ ] JOIN -- Joins a channel.
- [ ] PRIVMSG and NOTICE -- Sends a message to another nickname or channel.
- [ ] TOPIC -- Changes or returns the topic of a channel.
- [ ] NAMES -- Shows a list of which nicknames are on a channel.
- [ ] WHO -- Like NAMES, but returns more information. (Your client probably sends
  this when it joins a channel.)
- [ ] WHOIS -- Shows information about a nickname, including how long it has been
  connected.
- [ ] PART -- Leaves a channel.
- [ ] AWAY -- Marks or unmarks you as being away.
- [ ] INVITE -- Invites another user to a channel.
- [X] PING -- (Your client uses this command to measure the speed of its connection
  to the server.)
- [X] QUIT -- Disconnects from the server.
