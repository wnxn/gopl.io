package eval

import (
	"reflect"
	"testing"
)

func TestVar_String(t *testing.T) {
	tests := []struct{
		variable Var
		ex Expr
	}{
		{
			Var("x"),
			Var("x"),
		},
		{
			Var("x"),
			Var("x"),
		},
	}
	for _, test := range tests{
		expr, err := Parse(test.variable.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if !reflect.DeepEqual(test.ex, expr){
			t.Errorf("Parse %s: expect %v, but actually %v", test.variable.String(),
				test.ex,expr)
		}
	}
}

func TestLiteral_String(t *testing.T) {
	tests := []struct{
		l literal
	}{
		{
			literal(3.14345987978973897978979898973897978979898794),
		},
		{
			literal(23),
		},
	}
	for _, test := range tests{
		expr, err := Parse(test.l.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if !reflect.DeepEqual(test.l, expr){
			t.Errorf("Parse %s: expect %v, but actually %v", test.l.String(),
				test.l,expr)
		}
	}
}

func TestUnary_String(t *testing.T) {
	tests := []struct{
		u unary
	}{
		{
			unary{
				op: '+',
				x: literal(3.12321),
			},
		},
		{
			unary{
				op: '-',
				x: Var("asd"),
			},
		},
	}
	for _, test:=range tests{
		expr, err := Parse(test.u.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if test.u.String() != expr.String(){
			t.Errorf("Parse %s: expect %v, but actually %v", test.u.String(),
				test.u,expr)
		}
	}
}

func TestBinary_String(t *testing.T) {
	tests := []struct{
		b binary
	}{
		{
			binary{
				op: '*',
				x: unary{
					op: '-',
					x: Var("asd"),
				},
				y: Var("asd"),
			},
		},
		{
			binary{
				op: '+',
				x:binary{
					op: '*',
					x: unary{
						op: '-',
						x: Var("asd"),
					},
					y: Var("asd"),
				},
				y:binary{
					op:'/',
					x: literal(3.12321),
					y:  Var("asd"),
				},
			},
		},
	}
	for _, test:=range tests{
		expr, err := Parse(test.b.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if test.b.String() != expr.String(){
			t.Errorf("Parse %s: expect %v, but actually %v", test.b.String(),
				test.b,expr)
		}
	}
}

func TestCall_String(t *testing.T) {
	tests := []struct{
		c call
	}{
		{
			call{
				fn: "pow",
				args:[]Expr{
					binary{
						op: '*',
						x: unary{
							op: '-',
							x: Var("asd"),
						},
						y: Var("asd"),
					},
					literal(3.12321),
				},
			},
		},
	}
	for _, test:=range tests{
		expr, err := Parse(test.c.String())
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if test.c.String() != expr.String(){
			t.Errorf("Parse %s: expect %v, but actually %v", test.c.String(),
				test.c,expr)
		}
	}
}