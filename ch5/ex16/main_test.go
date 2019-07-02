package main

import (
	"testing"
)

func TestJoin1(t *testing.T){
	testcase := []struct{
		name string
		input []string
		output string
	}{
		{
			name: "case1",
			input: []string{
				"new",
				"star",
				"award",
			},
			output: "newstaraward",
		},
		{
			name: "case2",
			input: []string{},
			output:"",
		},
	}

	for _,v:=range testcase{
		res := join1(v.input...)
		if res != v.output{
			t.Errorf("name %s, expect %s but actually %s\n", v.name, v.output, res)
		}
	}
}

func TestJoin2(t *testing.T){
	testcase := []struct{
		name string
		input []string
		output string
	}{
		{
			name: "case1",
			input: []string{
				"new",
				"star",
				"award",
			},
			output: "newstaraward",
		},
		{
			name: "case2",
			input: []string{},
			output:"",
		},
	}

	for _,v:=range testcase{
		res := join2(v.input...)
		if res != v.output{
			t.Errorf("name %s, expect %s but actually %s\n", v.name, v.output, res)
		}
	}
}

func BenchmarkJoin1(b *testing.B){
	testcase := []struct{
		name string
		input []string
		output string
	}{
		{
			name: "case1",
			input: []string{
				"new",
				"star",
				"award",
			},
			output: "newstaraward",
		},
		{
			name: "case2",
			input: []string{},
			output:"",
		},
		{
			name: "case3",
			input:[]string{
				"new",
				"star",
				"award",
				"asdasedfcsfsefefe",
				"weq3wed3ferfw4rr3dw2e2e22dd",
			},
			output:"newstarawardasdasedfcsfsefefeweq3wed3ferfw4rr3dw2e2e22dd",
		},
	}
	for _,v:=range testcase{
		for i:=0;i< b.N;i++{
			res := join1(v.input...)
			if res != v.output{
				b.Errorf("name %s, expect %s but actually %s\n", v.name, v.output, res)
			}
		}
	}
}
// strings.jon 139 ns/op
// join1 277 ns/op
// join2 339 ns/op
// join3 210 ns/op
func BenchmarkJoin2(b *testing.B){
	testcase := []struct{
		name string
		input []string
		output string
	}{
		{
			name: "case1",
			input: []string{
				"new",
				"star",
				"award",
			},
			output: "newstaraward",
		},
		{
			name: "case2",
			input: []string{},
			output:"",
		},
		{
			name: "case3",
			input:[]string{
				"new",
				"star",
				"award",
				"asdasedfcsfsefefe",
				"weq3wed3ferfw4rr3dw2e2e22dd",
			},
			output:"newstarawardasdasedfcsfsefefeweq3wed3ferfw4rr3dw2e2e22dd",
		},
	}
	for _,v:=range testcase{
		for i:=0;i< b.N;i++{
			res := join3(v.input...)
			if res != v.output{
				b.Errorf("name %s, expect %s but actually %s\n", v.name, v.output, res)
			}
		}
	}
}