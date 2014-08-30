package message

import (
	"strings"
)

// prefix =  servername / ( nickname [ [ "!" user ] "@" host ] )
type prefix struct {
	name string
	user string
	host string
}

func Prefix(args ...string) *prefix {
	switch len(args) {
	case 1:
		return &prefix{args[0], "", ""}
	case 2:
		return &prefix{args[0], "", args[1]}
	case 3:
		return &prefix{args[0], args[1], args[2]}
	default:
		return &prefix{"", "", ""}
	}
}

func (p *prefix) String() string {
	s := p.name

	if p.user != "" {
		s += "!" + p.user
	}

	if p.host != "" {
		s += "@" + p.host
	}

	return s
}

func parsePrefix(s string) *prefix {
	parts := strings.Split(s, "@")
	if len(parts) == 1 {
		return Prefix(parts...)
	}

	left := strings.Split(parts[0], "!")
	if len(left) == 1 {
		return Prefix(left[0], parts[1])
	}

	return Prefix(left[0], left[1], parts[1])
}
