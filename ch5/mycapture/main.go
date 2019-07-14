package main

import "fmt"

func main() {
	capture1()
}

func capture1() {
	dirs := []string{"test1", "test2", "test3"}
	res := []func(){}
	for _, val := range dirs {
		fmt.Printf("%s %0xd\n", val, &val)
		res = append(res, func() {
			fmt.Printf("%s %0xd\n", val, &val)
		})
	}
	for i := range res {
		res[i]()
	}
}
