package ex1

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		str     string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			str:     "asdwdewfrf",
			counts:  map[rune]int{'a': 1, 'd': 2, 'e': 1, 'f': 2, 's': 1, 'w': 2, 'r': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 10},
			invalid: 0,
		},
		{
			str:     "。27世 15 界 14 小于等于",
			counts:  map[rune]int{'7': 1, '5': 1, '4': 1, '小': 1, '2': 1, '世': 1, ' ': 4, '1': 2, '界': 1, '于': 2, '等': 1, '。': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 10, 0, 7},
			invalid: 0,
		},
	}
	for _, test := range tests {
		resCount, resUtflen, resInvalid := charcount(strings.NewReader(test.str))
		if !reflect.DeepEqual(resCount, test.counts) {
			t.Errorf("charcount(%s)= expect %v, actually %v", test.str, test.counts, resCount)
		}
		if !reflect.DeepEqual(resUtflen, test.utflen) {
			t.Errorf("charcount(%s)= expect %v, actually %v", test.str, test.utflen, resUtflen)
		}
		if resInvalid != test.invalid {
			t.Errorf("charcount(%s)= expect %v, actually %v", test.str, test.invalid, resInvalid)
		}
	}
}
