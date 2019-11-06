package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// reflect.Type
	i := 3
	t := reflect.TypeOf(i)
	fmt.Println(t.String()) // int
	fmt.Println(t)          // int
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File

	fmt.Printf("%T\n", 3)

	// reflect.Value
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String(), v.Kind())
}
