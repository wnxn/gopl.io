package main

import "testing"

func TestTrans(t *testing.T) {
	testcases := []struct {
		input  string
		output string
	}{
		{
			input:  "$foo",
			output: "gpp",
		},
		{
			input:  "$$foo",
			output: "hqq",
		},
		{
			input:  "$",
			output: "",
		},
		{
			input:  "$$",
			output: "",
		},
	}
	for _, v := range testcases {
		res := Trans(v.input)
		if res != v.output {
			t.Errorf("input %s, expect %s, but actually %s\n", v.input, v.output, res)
		}
	}
}

func TestExpand(t *testing.T) {
	testcases := []struct {
		input    string
		function func(string) string
		output   string
	}{
		{
			input:    "dwad$$foo$qwe",
			function: Trans,
			output:   "dwadhqqrxf",
		},
		{
			input:    "dwad$foo",
			function: Trans,
			output:   "dwadgpp",
		},
	}
	for _, v := range testcases {
		res := Expand(v.input, v.function)
		if res != v.output {
			t.Errorf("input %s, output %s, but actually %s\n", v.input, v.output, res)
		}
	}
}
