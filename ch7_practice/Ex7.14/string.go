package eval

import "fmt"

func (v Var) String() string {
	return fmt.Sprintf(string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%s", u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("( %s %c %s )", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	var str string
	str = fmt.Sprintf("%s(", c.fn)

	for i, arg := range c.args {
		if i > 0 {
			str += fmt.Sprintf(",")
		}
		str += arg.String()
	}

	str += fmt.Sprintf(")")
	return str
}
