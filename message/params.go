package message

// params     =  *14( SPACE middle ) [ SPACE ":" trailing ]
//            =/ 14( SPACE middle ) [ SPACE [ ":" ] trailing ]
type params struct {
	l []string
	t string
}

func Params(args []string) *params {
	return &params{args, ""}
}

func ParamsT(args []string, trailing string) *params {
	return &params{args, trailing}
}

func (ps *params) String() string {
	s := ""

	for _, p := range ps.l {
		s += " " + p
	}

	if ps.t != "" {
		s += " :" + ps.t
	}

	return s
}

func (ps *params) all() []string {
	if ps.t != "" {
		return append(ps.l, ps.t)
	}

	return ps.l
}

func (ps *params) Any() bool {
	return ps.Len() != 0
}

func (ps *params) Len() int {
	return len(ps.all())
}

func (ps *params) Get(i int) string {
	return ps.all()[i]
}
