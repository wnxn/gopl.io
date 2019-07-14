package main

import "fmt"

func main() {
	handleRecover()
}

func handleRecover() {
	defer func() {
		p := recover()
		fmt.Println(p)
	}()
	nonZeroValue()
}

func nonZeroValue() {
	str := "nonZero"
	panic(str)

}
