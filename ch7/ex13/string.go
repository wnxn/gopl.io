package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return fmt.Sprintf("%s", u.x.String())
	case '-':
		return fmt.Sprintf("(-%s)", u.x.String())
	default:
		return "invalid"
	}
}

func (b binary) String() string {
	return fmt.Sprintf("(%s%c%s)", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	str := &bytes.Buffer{}
	fmt.Fprintf(str, "%s(", c.fn)
	for i, k := range c.args {
		if i > 0 {
			fmt.Fprint(str, ",")
		}
		fmt.Fprintf(str, "%s", k.String())
	}
	fmt.Fprint(str, ")")
	return str.String()
}
