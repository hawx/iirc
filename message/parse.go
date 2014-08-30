package message

import (
	"strings"
)

func Parse(line string) M {
	pr, co, ar, tr := split(line)

	var prfx *prefix
	if pr != "" {
		prfx = parsePrefix(pr)
	}

	command := co

	var prms *params
	if tr == "" {
		prms = Params(ar)
	} else {
		prms = ParamsT(ar, tr)
	}

	return Message3(prfx, command, prms)
}

func split(line string) (string, string, []string, string) {
	s := strings.TrimRight(line, "\r\n")

	prefix := ""
	trailing := ""
	var args []string

	if s[0] == ':' {
		parts := strings.SplitN(s[1:], " ", 2)
		prefix, s = parts[0], parts[1]
	}

	if strings.Index(s, " :") != -1 {
		parts := strings.SplitN(s, " :", 2)
		s, trailing = parts[0], parts[1]
		args = strings.Split(s, " ")
	} else {
		args = strings.Split(s, " ")
	}

	command := args[0]
	args = args[1:]

	return prefix, command, args, trailing
}
