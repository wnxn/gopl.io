package ex6

import "testing"

func TestPopCountTable(t *testing.T) {
	tests := []struct {
		x      uint64
		result int
	}{
		{
			x:      0,
			result: 0,
		},
		{
			x:      0xABCDEF,
			result: 17,
		},
	}
	for _, test := range tests {
		res := PopCountTable(test.x)
		if res != test.result {
			t.Errorf("PopCountTable(%d)= expect %d, but actually %d", test.x, test.result, res)
		}
	}
}

func TestPopCountMove(t *testing.T) {
	tests := []struct {
		x      uint64
		result int
	}{
		{
			x:      0,
			result: 0,
		},
		{
			x:      0xABCDEF,
			result: 17,
		},
	}
	for _, test := range tests {
		res := PopCountMove(test.x)
		if res != test.result {
			t.Errorf("PopCountMove(%d)= expect %d, but actually %d", test.x, test.result, res)
		}
	}
}

func TestPopCountClearRight(t *testing.T) {
	tests := []struct {
		x      uint64
		result int
	}{
		{
			x:      0,
			result: 0,
		},
		{
			x:      0xABCDEF,
			result: 17,
		},
	}
	for _, test := range tests {
		res := PopCountClearRight(test.x)
		if res != test.result {
			t.Errorf("PopCountClearRight(%d)= expect %d, but actually %d", test.x, test.result, res)
		}
	}
}

// BenchmarkPopCountTable-8   	2'000'000'000	         0.29 ns/op
func BenchmarkPopCountTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTable(uint64(i % 64))
	}
}

// BenchmarkPopCountMove-8   	300'000'000	         5.83 ns/op
func BenchmarkPopCountMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountMove(uint64(i % 64))
	}
}

// BenchmarkPopCountClearRight-8   	500`000`000	         3.85 ns/op
func BenchmarkPopCountClearRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearRight(uint64(i % 64))
	}
}
