package memsync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMemsync(t *testing.T) {
	Memsync()

	time.Sleep(time.Second)
}

var done = make(chan struct{})

func cancel() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func TestRoutine(t *testing.T) {
	fmt.Println(cancel())
	fmt.Println(cancel())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()
	fmt.Println(cancel())
	fmt.Println(cancel())

}
