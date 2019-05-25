package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var types  = flag.String("types", "sha256", "support sha256, sha384, sha512")

func main() {
	flag.Parse()
	str := ""
	fmt.Scanln(&str)
	fmt.Printf("%x\n", handler(types,str))
}

func handler(types *string, str string)string{
	var bytes []byte
	switch *types {
	case "sha256":
		for _,v:=range sha256.Sum256([]byte(str)){
			bytes = append(bytes,v)
		}
	case "sha384":
		for _,v:=range sha512.Sum384([]byte(str)){
			bytes = append(bytes,v)
		}
	case "sha512":
		for _,v:=range sha512.Sum512([]byte(str)){
			bytes=append(bytes,v)
		}
	}
	return string(bytes)
}