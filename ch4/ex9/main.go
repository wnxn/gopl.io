package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	b, _,_:=in.ReadLine()
	wordfreq(b)
	for k,v:=range wordfreq(b){
		fmt.Println(k,v)
	}
}

func wordfreq(b []byte)map[string]int{
	btext := b
	res := make(map[string]int)
	for len(btext)>0{
		adv, token, _:=bufio.ScanWords(btext,true)
		btext=btext[adv:]
		res[string(getWord(token))]++
	}
	return res
}

func getWord(b []byte)[]byte{
	res := b[:0]
	for _,v:=range b{
		if unicode.IsLetter(rune(v)){
			res = append(res,v)
		}
	}
	return res
}