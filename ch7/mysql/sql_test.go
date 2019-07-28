package mysql

import (
	"testing"
)

func TestSqlQuote1(t *testing.T) {
	tests := []struct{
		x interface{}
		res string
	}{
		{
			nil,
			"NULL",
		},
		{
			23,
			"23",
		},
		{
			true,
			"TRUE",
		},
		{
			"hello",
			"hello",
		},
	}
	for _,v:=range tests{
		str := SqlQuote2(v.x)
		if str != v.res{
			t.Errorf("expect %s, but actually %s", v.res, str)
		}
	}
}