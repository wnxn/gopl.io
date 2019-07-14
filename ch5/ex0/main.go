package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	//readFromBytes([]byte("hello,世界！"))
	double(4)
}

func readFromBytes(str []byte) []rune {
	res := []rune{}
	reader := strings.NewReader(string(str))
	buf := bufio.NewReader(reader)
	for {
		r, s, err := buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				return res
			}
			fmt.Errorf("%v\n", err)
			return res
		}
		fmt.Printf("get rune %q size %d\n", r, s)
		res = append(res, r)
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}
