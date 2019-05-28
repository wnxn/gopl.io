package main

import "fmt"

func reverse(arr *[5]int)*[5]int{
	for i,j:=0,len(arr)-1;i<j;i,j=i+1,j-1{
		arr[i],arr[j]=arr[j],arr[i]
	}
	return arr
}

func main() {
	arr := [...]int{2,3,5,3,6}
	fmt.Println(reverse(&arr))
}
