package main

import (
	"fmt"
)

func main() {
	res := min(2,3,4)
	fmt.Println(res)
}

func max(first float64, args ...float64)float64{
	res := first
	for _,val:=range args{
		if val > res{
			res = val
		}
	}
	return res
}

func min(first float64, args ...float64)float64{
	res := first
	for _,val:=range args{
		if val < res{
			res = val
		}
	}
	return res
}