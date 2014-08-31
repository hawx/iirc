package handler

import "github.com/hawx/iirc/reply"

func PrivMsg(c Client, s Server, args []string) {
	subject, ok := s.Find(args[0])
	if ok {
		msg := reply.PrivMsg(c.Name(), c.UserName(), s.Name(), args[0], args[1])
		subject.SendExcept(msg, c.Name())
	}
}
