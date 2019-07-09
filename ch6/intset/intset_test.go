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
