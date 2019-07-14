package main

import (
	"bytes"
	"fmt"
	"io"
)

type myreader struct {
	r io.Reader
	n int
}

func main() {
	str := "hello world"
	var r1 io.Reader
	r1 = bytes.NewReader([]byte(str))
	r2 := LimitReader(r1, 5)
	var con2 []byte = make([]byte, 2)
	n, _ := r2.Read(con2)
	fmt.Printf("%d, %s\n", n, string(con2))
	var con3 []byte = make([]byte, 3)
	n, _ = r2.Read(con3)
	fmt.Printf("%d, %s\n", n, string(con3))
}

func (read *myreader) Read(p []byte) (int, error) {
	if len(p) <= 0 {
		return 0, io.EOF
	}
	if len(p) > read.n {
		p = p[:read.n]
	}
	sub, err := read.r.Read(p)
	if err != nil {
		return 0, err
	}
	read.n -= sub
	return sub, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &myreader{r, int(n)}
}
