package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var a float32 = 3
	//b  := 4.3
	fmt.Println(unsafe.Sizeof(a))
}
