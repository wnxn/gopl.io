package ex4

import (
	"fmt"
)

func Pipeline(num int){
	var chs []chan string
	for i:= 0;i<num;i++{
		chs = append(chs, make(chan string))
	}
	go func(){
		chs[0]<- "hello"
	}()
	for i:=1;i<num;i++{
		go passValue(chs[i-1],chs[i])
	}
	fmt.Println(len(<-chs[num-1]))
}

func passValue(ch1 <-chan string, ch2 chan<- string){
	str:=<-ch1
	ch2 <- str +"p"
}