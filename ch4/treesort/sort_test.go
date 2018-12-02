// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/adonovan/gopl.io/ch4/treesort"
	"fmt"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	treesort.Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestString(t *testing.T){
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	tree := treesort.GetTree(data)
	fmt.Println(tree)
}
