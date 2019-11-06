package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Println(reflect.TypeOf(w).Kind())

	name := []string{"ws", "fe"}
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(name).Kind())
	switch interface{}(name).(type) {
	case []string:
		fmt.Println("[]string")
	default:
		fmt.Println("unknown")
	}
}
