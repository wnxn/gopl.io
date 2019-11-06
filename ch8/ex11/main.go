package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var done = make(chan struct{})

func cancel() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("https://sh.sina.com.cn") }()
	go func() { responses <- request("https://zj.sina.com.cn") }()
	go func() { responses <- request("https://ah.sina.com.cn") }()

	return <-responses
}

func request(hostname string) (response string) {

	fmt.Println("start" + hostname)
	req, err := http.NewRequest(http.MethodGet, hostname, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if cancel() {
		panic("hi")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	bytes := bytes.Buffer{}
	err = resp.Header.Write(&bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("end" + hostname)
	return "parsing " + hostname + bytes.String()
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	fmt.Println(mirroredQuery())
}
