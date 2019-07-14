package main

import "fmt"

func main() {
	var a1 []int
	printResult(a1)
	var a2 []int = nil
	printResult(a2)
	a3 := []int(nil)
	printResult(a3)
	a4 := []int{}
	printResult(a4)
	a6 := make([]int, 0)
	printResult(a6)
	a7 := make([]int, 0)[:]
	printResult(a7)

	a8 := make([]int, 9)
	printResult(a8)
	a9 := make([]int, 9, 11)
	printResult(a9)
	a10 := make([]int, 11)[:9]
	printResult(a10)
	sliceLenAdditional()
}

func printResult(s []int) {
	fmt.Printf("%d, %d, %#v\n", cap(s), len(s), 9)
}

func sliceLenAdditional() {
	a := make([]int, 9, 11)
	for i := len(a); i < 100; i++ {
		a = append(a, 1)
		printResult(a)
	}
}
