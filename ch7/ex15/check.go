// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"fmt"
	"strings"
)

//!+Check

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) GetExpr() []Expr {
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) GetExpr() []Expr {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) GetExpr() []Expr {
	return []Expr{u.x}
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) GetExpr() []Expr {
	return []Expr{b.x, b.y}
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) GetExpr() []Expr {
	return c.args
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

//!-Check

func (m min) Check(vars map[Var]bool) error {
	arity, ok := minParams[m.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", m.fn)
	}
	if len(m.args) < arity {
		return fmt.Errorf("call to %s has %d args, want >= %d",
			m.fn, len(m.args), arity)
	}
	for _, arg := range m.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (m min) GetExpr() []Expr {
	return m.args
}

var minParams = map[string]int{"min": 1}
