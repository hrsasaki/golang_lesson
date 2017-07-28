package eval

import (
	"strconv"
	"strings"
)

//!+String

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'g', 4, 64)
}

func (u unary) String() string {
	return string(u.op) + "[" + u.x.String() + "]"
}

func (b binary) String() string {
	return string(b.op) + "[" + b.x.String() + ", " + b.y.String() + "]"
}

func (c call) String() string {
	var argstrs []string
	for _, arg := range c.args {
		argstrs = append(argstrs, arg.String())
	}
	args := strings.Join(argstrs, ", ")
	return c.fn + "[" + args + "]"
}

//!-String
