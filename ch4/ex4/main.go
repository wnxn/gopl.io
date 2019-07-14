package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(s, 3))
}

func rotate(s []int, p int) []int {
	if p > len(s) || p < 1 {
		return s
	}
	s = append(s, s[:p]...)
	return s[p:]
}
