package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	temp := &bytes.Buffer{}
	w, c := CountingWriter(temp)
	fmt.Fprint(w, "hello world")
	fmt.Fprint(w, "hello world")
	fmt.Println(*c)

}

type MyWriter struct {
	w io.Writer
	c int64
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	res := MyWriter{
		w: w,
		c: 0,
	}
	return &res, &res.c
}

func (w *MyWriter) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	if err != nil {
		return 0, err
	}
	w.c += int64(n)
	return int(w.c), nil
}
