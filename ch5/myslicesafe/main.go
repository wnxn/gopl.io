package main

import "fmt"

func modifySlice(s []string){
	s[1]="sf"
	fmt.Printf("content=%s, address=%p, first=%p\n",s, &s, &(s)[0])
}

func main() {
	s := make([]string, 9)
	s =[]string{"hello","waf"}
	fmt.Printf("content=%s, address=%p, first=%p\n",s, &s, &(s)[0])
	modifySlice(s)
	fmt.Printf("content=%s, address=%p, first=%p\n",s, &s, &(s)[0])
}
