package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int

type Lock struct {
	ch chan int
}

func New() Lock {
	return Lock{
		ch: make(chan int, 1),
	}
}

func (l *Lock) Lock() {
	<-l.ch
}

func (l *Lock) UnLock() {
	l.ch <- 1
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100000)
	t := New()
	t.ch <- 1

	for i := 0; i < 100000; i++ {
		go func(j int) {
			t.Lock()
			defer t.UnLock()
			counter++
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(counter)
}
