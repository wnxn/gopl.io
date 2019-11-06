package ex5

import (
	"fmt"
	"time"
)

func Pingpong() {
	ch1, ch2 := make(chan int), make(chan int)
	timer := time.Tick(1 * time.Second)
	go func() {
		ch1 <- 1
	}()
	go func() {
		Man(ch1, ch2)
	}()
	go func() {
		Man(ch2, ch1)
	}()
	<-timer
	close(ch1)
	close(ch2)
	fmt.Println(<-ch1, <-ch2)
}

func Man(in <-chan int, out chan<- int) {
	for {
		select {
		case cnt := <-in:
			if cnt == 0 {
				return
			}
			cnt++
			out <- cnt
		}
	}
}
