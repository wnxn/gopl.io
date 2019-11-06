package memsync

import "fmt"

func Memsync() {
	var x, y int
	go func() {
		x = 1
		fmt.Println("y = ", y, " ")
	}()
	go func() {
		y = 2
		fmt.Println("x = ", x, " ")
	}()
}
