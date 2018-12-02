package main

import (
	"io"
	"fmt"
	"bytes"
	"github.com/golang/glog"
)

type MyWriter struct{
	cnt int64
	w io.Writer
}

func CountingWriter(w io.Writer)(io.Writer, *int64){
	res := MyWriter{0, w}
	return &res, &res.cnt
}

func (s *MyWriter)Write(p []byte)(int, error){
	num, err := s.w.Write(p)
	if err == nil{
		s.cnt += int64(num)
	}
	return num, err
}

func main() {
	temp := &bytes.Buffer{}
	temp2, cnt := CountingWriter(temp)
	fmt.Fprintf(temp2, "hello")
	fmt.Fprintf(temp2, "hello")
	glog.Info(*cnt)
}
