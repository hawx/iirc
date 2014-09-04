package message

import (
	"strings"
)

// message =  [ ":" prefix SPACE ] command [ params ] crlf
type M struct {
	Prefix  *prefix
	Command string
	Params  *params
}

func Message3(Prefix *prefix, command string, prms *params) M {
	return M{Prefix, command, prms}
}

func (m M) Args() []string {
	return m.Params.all()
}

func (m M) String() string {
	s := ""
	if m.Prefix != nil {
		s += ":" + m.Prefix.String() + " "
	}

	s += m.Command

	if m.Params != nil {
		s += m.Params.String()
	}

	return s + "\r\n"
}

func (m M) Parts() []M {
	trailing := m.Params.t
	m.Params.t = ""

	maxLength := 512 - len(m.String())
	splits := []string{}
	currSplit := ""

	for _, c := range strings.Split(trailing, " ") {
		if len(c)+len(currSplit)+1 <= maxLength {
			currSplit += " " + c
		} else {
			splits = append(splits, currSplit[1:])
			currSplit = " " + c
		}
	}

	if currSplit != "" {
		splits = append(splits, currSplit[1:])
	}

	parts := []M{}
	for _, split := range splits {
		prms := &params{m.Params.l, split}
		part := M{m.Prefix, m.Command, prms}
		parts = append(parts, part)
	}

	return parts
}
