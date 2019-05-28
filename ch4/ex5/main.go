package main

import "fmt"

func main() {
	s := []string{"wa","ha","ha","like","me","ha"}
	fmt.Println(eliminate(s))
}

func eliminate(s []string)[]string{
	out := s[:0]
	var prev string = ""
	for _,v:=range s{
		if v != prev{
			out = append(out, v)
		}
		prev=v
	}
	return out
}