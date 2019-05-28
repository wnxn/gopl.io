package main

import "fmt"

func main() {
	s :="wangxin"
	fmt.Println(string(reverse([]byte(s))))
}

func reverse(s []byte)[]byte{
	return  help(s,0,len(s)-1)
}

func help(s[]byte, i,j int)[]byte{
	if i < j{
		s[i],s[j]=s[j],s[i]
		help(s,i+1,j-1)
	}
		return s
}