package main

import (
	"io"
	"strings"
	"fmt"
)

type limitReader struct{
	reader io.Reader
	num int
}

func (r *limitReader)Read(p []byte)(int, error){
	if r.num <= 0{
		return 0, io.EOF
	}
	if len(p) >r.num{
		p = p[:r.num]
	}
	n, err := r.reader.Read(p)
	r.num -= n
	return n,err
}

func LimitReader(r io.Reader, n int64)io.Reader{
	return &limitReader{r, int(n)}
}

func main() {
	testString := "Now is the winter of our discontent"
	var strRdr io.Reader = strings.NewReader(testString)

	reader := LimitReader(strRdr, 5)
	a := make([]byte, 3)
	n, err := reader.Read(a)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(n, string(a),a)

	b := make([]byte, 4)
	n, err = reader.Read(b)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(n, string(b),b)

	c := make([]byte, 5)
	n, err = reader.Read(c)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(n, string(c),c)
}
