package ex5

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		str string
		sep string
		res []string
	}{
		{
			str: "hello,world",
			sep: "o",
			res: []string{"hell", ",w", "rld"},
		},
	}
	for _, test := range tests {
		res := strings.Split(test.str, test.sep)
		if !reflect.DeepEqual(test.res, res) {
			t.Errorf("Split(%s,%s)=%s, want %s", test.str, test.sep, res, test.res)
		}
	}
}
