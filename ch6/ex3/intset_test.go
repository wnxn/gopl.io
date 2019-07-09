// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"testing"
)

func TestExample_one(t *testing.T) {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	t.Log(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	t.Log(y.String()) // "{9 42}"

	x.UnionWith(&y)
	t.Log(x.String()) // "{1 9 42 144}"

	t.Log(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func TestExample_two(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	t.Log(&x)         // "{1 9 42 144}"
	t.Log(x.String()) // "{1 9 42 144}"
	t.Log(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestIntSet_Len(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if res :=x.Len();res != 4{
		t.Errorf("Len expect 4, but actually %d",res)
	}

}

func TestIntSet_Remove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if res :=x.Has(9); res != true{
		t.Errorf("expect true, but actually %t",res)
	}
	x.Remove(9)
	if res :=x.Has(9); res != false{
		t.Errorf("expect false, but actually %t",res)
	}
}

func TestIntSet_Clear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if res :=x.Len(); res != 4{
		t.Errorf("expect true, but actually %d",res)
	}
	x.Clear()
	if res :=x.Len(); res != 0{
		t.Errorf("expect false, but actually %d",res)
	}
}

func TestIntSet_Copy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	if res :=x.Len(); res != 4{
		t.Errorf("expect 4, but actually %d",res)
	}

	y:=x.Copy()
	if res :=y.Len(); res != 4{
		t.Errorf("expect 4, but actually %d",res)
	}

	x.Clear()
	if res :=x.Len(); res != 0{
		t.Errorf("expect 0, but actually %d",res)
	}
	if res :=y.Len(); res != 4{
		t.Errorf("expect 4, but actually %d",res)
	}
}

func TestIntSet_AddAll(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)
	x.AddAll(9,10,34)
	if res :=x.Len(); res != 6{
		t.Errorf("expect 6, but actually %d",res)
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	var y IntSet
	y.Add(1213)
	y.Add(1424)
	y.Add(9)
	y.Add(422)

	x.DifferenceWith(&y)
	if res :=x.Len(); res != 6{
		t.Errorf("expect 6, but actually %d",res)
	}
}