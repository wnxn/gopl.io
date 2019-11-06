package intset

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

//func TestIntSet_Add(t *testing.T) {
//	intset := IntSet{}
//
//	tests := []struct{
//		add int
//		str string
//	}{
//		{
//			add: 9,
//			str: "{9}",
//		},
//		{
//			add: 65,
//			str: "{9 65}",
//		},
//		{
//			add: 12,
//			str: "{9 12 65}",
//		},
//	}
//	for _, test := range tests{
//		intset.Add(test.add)
//		if test.str != intset.String(){
//			t.Errorf("IntSet.Add(%d) =  expect %s, but actually %s", test.add, test.str, intset.String())
//		}
//	}
//}
//
//func TestIntSet_Has(t *testing.T) {
//	intset := IntSet{
//		words:[]uint64{0x1000000},
//	}
//	tests := []struct{
//		num int
//		has bool
//	}{
//		{
//			num:21,
//			has: false,
//		},
//		{
//			num:24,
//			has:true,
//		},
//	}
//	for _, test := range tests{
//		res:= intset.Has(test.num)
//		if res != test.has{
//			t.Errorf("IntSet.Has(%d) = expect %t, but actually %t", test.num, test.has, res)
//		}
//	}
//}
//
//func TestIntSet_UnionWith(t *testing.T) {
//	tests := []struct{
//		set1 IntSet
//		set2 IntSet
//		res IntSet
//	}{
//		{
//			set1: IntSet{
//				words:[]uint64{0x1003000},
//			},
//			set2: IntSet{
//				words:[]uint64{0x000F000, 0x1},
//			},
//			res: IntSet{
//				words:[]uint64{0x100F000,0x1},
//			},
//		},
//	}
//	for _, test := range tests {
//		test.set1.UnionWith(&test.set2)
//		if test.set1.String() != test.res.String() {
//			t.Errorf("IntSet.Union(%s) = expect %s, but actually %s", test.set1.String(), test.set2.String(),
//				test.res.String())
//		}
//	}
//}

func TestIntSetMap_UnionWith(t *testing.T) {
	tests := []struct {
		set1 *IntSetMap
		set2 *IntSetMap
		res  *IntSetMap
	}{
		{
			set1: NewIntSetMap(),
			set2: NewIntSetMap(),
			res:  NewIntSetMap(),
		},
	}
	for _, test := range tests {
		test.set1.Add(23)
		test.set1.Add(45)
		test.set2.Add(45)
		test.set2.Add(56)
		test.res.Add(23)
		test.res.Add(45)
		test.res.Add(56)
		test.set1.UnionWith(test.set2)
		if !reflect.DeepEqual(test.res.mp, test.set1.mp) {
			t.Errorf("expect %v, but actually %v", test.res.mp, test.set1.mp)
		}
	}
}

// BenchmarkIntSet_Add-8   	100`000`000	        16.1 ns/op
func BenchmarkIntSet_Add(b *testing.B) {
	set := IntSet{}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(0xFF))
	}
}

// BenchmarkIntSet_Has-8   	100`000`000	        12.4 ns/op
func BenchmarkIntSet_Has(b *testing.B) {
	set := IntSet{}
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		set.Has(rng.Intn(0xFF))
	}
}

// BenchmarkIntSet_UnionWith-8   	500`000`000	         3.64 ns/op
func BenchmarkIntSet_UnionWith(b *testing.B) {
	set1 := &IntSet{}
	set2 := &IntSet{}
	set1.Add(23)
	set1.Add(45)
	set2.Add(45)
	set2.Add(56)
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

// BenchmarkIntSetMap_Add-8   	30`000`000	        42.7 ns/op
func BenchmarkIntSetMap_Add(b *testing.B) {
	set := NewIntSetMap()
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(0xFF))
	}
}

// BenchmarkIntSetMap_Has-8   	30`000`000	        44.3 ns/op
func BenchmarkIntSetMap_Has(b *testing.B) {
	set := NewIntSetMap()
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		set.Add(rng.Intn(0xFF))
	}
}

// BenchmarkIntSetMap_UnionWith-8   	20`000`000	        61.8 ns/op
func BenchmarkIntSetMap_UnionWith(b *testing.B) {
	set1 := NewIntSetMap()
	set2 := NewIntSetMap()
	set1.Add(23)
	set1.Add(45)
	set2.Add(45)
	set2.Add(56)
	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}
