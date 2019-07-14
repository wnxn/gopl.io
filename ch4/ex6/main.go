package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "wad\n\tsfds\tfsd"
	fmt.Printf("%v\n", string([]byte(s)))
	fmt.Printf("%v\n", string(squash([]byte(s))))
}

func squash(b []byte) []byte {
	out := b[:0]
	flag := false
	for _, v := range b {
		if flag == false {
			out = append(out, v)
		} else {
			out[len(out)-1] = ' '
		}
		if unicode.IsSpace(rune(v)) {
			flag = true
		} else {
			flag = false
		}

	}
	return out
}
