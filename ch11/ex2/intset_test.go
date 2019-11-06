package intset

import "testing"

func TestIntSet_Add(t *testing.T) {
	intset := IntSet{}

	tests := []struct {
		add int
		str string
	}{
		{
			add: 9,
			str: "{9}",
		},
		{
			add: 65,
			str: "{9 65}",
		},
		{
			add: 12,
			str: "{9 12 65}",
		},
	}
	for _, test := range tests {
		intset.Add(test.add)
		if test.str != intset.String() {
			t.Errorf("IntSet.Add(%d) =  expect %s, but actually %s", test.add, test.str, intset.String())
		}
	}
}

func TestIntSet_Has(t *testing.T) {
	intset := IntSet{
		words: []uint64{0x1000000},
	}
	tests := []struct {
		num int
		has bool
	}{
		{
			num: 21,
			has: false,
		},
		{
			num: 24,
			has: true,
		},
	}
	for _, test := range tests {
		res := intset.Has(test.num)
		if res != test.has {
			t.Errorf("IntSet.Has(%d) = expect %t, but actually %t", test.num, test.has, res)
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	tests := []struct {
		set1 IntSet
		set2 IntSet
		res  IntSet
	}{
		{
			set1: IntSet{
				words: []uint64{0x1003000},
			},
			set2: IntSet{
				words: []uint64{0x000F000, 0x1},
			},
			res: IntSet{
				words: []uint64{0x100F000, 0x1},
			},
		},
	}
	for _, test := range tests {
		test.set1.UnionWith(&test.set2)
		if test.set1.String() != test.res.String() {
			t.Errorf("IntSet.Union(%s) = expect %s, but actually %s", test.set1.String(), test.set2.String(),
				test.res.String())
		}
	}
}
