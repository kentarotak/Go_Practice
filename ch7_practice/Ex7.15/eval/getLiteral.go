// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import "fmt"

// Format formats an expression as a string.
// It does not attempt to remove unnecessary parens.

func GetLiteral(buf map[Var]float64, e Expr) {
	switch e := e.(type) {
	case literal:
		//fmt.Fprintf(buf, "%g", e)

	case Var:
		buf[e] = 0

	case unary:
		//fmt.Fprintf(buf, "(%c", e.op)
		GetLiteral(buf, e.x)
		//buf.WriteByte(')')

	case binary:
		//buf.WriteByte('(')
		GetLiteral(buf, e.x)
		//fmt.Fprintf(buf, " %c ", e.op)
		GetLiteral(buf, e.y)
		//buf.WriteByte(')')

	case call:
		//fmt.Fprintf(buf, "%s(", e.fn)
		for i, arg := range e.args {
			if i > 0 {
				//	buf.WriteString(", ")
			}
			GetLiteral(buf, arg)
		}
		//buf.WriteByte(')')

	default:
		panic(fmt.Sprintf("unknown Expr: %T", e))
	}
}
