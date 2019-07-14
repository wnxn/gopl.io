package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, value := range os.Args {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}
