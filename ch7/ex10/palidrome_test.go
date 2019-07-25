package ex10

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct{
		name string
		input sort.Interface
		output bool
	}{
		{
			name:"normal",
			input: &sort.IntSlice{1,2,3,5,3,2,1},
			output:true,
		},
		{
			"zero value",
			&sort.StringSlice{},
			true,
		},
		{
			"failed ints",
			&sort.IntSlice{1,2,3,5,3,2},
			false,
		},
	}
	for _, v:=range testcases{
		res:=IsPalindrome(v.input)
		if res != v.output{
			t.Errorf("name %s: expect %t, but actually %t", v.name, v.output, res)
		}
	}
}