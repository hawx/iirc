package message

type Sender interface {
	Name() string
	Send(M)
	SendExcept(M, string)
}
