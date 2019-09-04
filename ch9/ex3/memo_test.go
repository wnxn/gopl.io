// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"sync"
	"testing"
	"github.com/adonovan/gopl.io/ch9/ex3"
	"gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	wg := sync.WaitGroup{}
	go func(){
		m.Cancel()
		wg.Done()
	}()
	memotest.Concurrent(t, m)
	wg.Wait()
}
