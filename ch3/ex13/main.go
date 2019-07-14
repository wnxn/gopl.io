package main

import "fmt"

const (
	x  = 1000
	KB = x
	MB = KB * x
	GB = MB * x
)

func main() {
	//fmt.Println(GB)
	var f float64 = 1e23
	fmt.Println(f)
}
