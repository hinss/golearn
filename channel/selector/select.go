package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generate a chan of int that will send out number n within 1500 Millisecond
func generator() chan int {

	c := make(chan int)

	go func() {
		n := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <- n
			n++
		}
	}()

	return c
}

func worker(id int, c chan int) {
	for {
		time.Sleep(time.Second)
		fmt.Printf("Workder:%d received %d\n",
			id, <-c)
	}
}

// channel 作为一等公民也可以作为函数的返回值、参数传递
func createWorker(i int) chan<- int {
	c := make(chan int)
	go worker(i, c)
	return c
}

func main() {

	c1, c2 := generator(), generator()
	var values []int
	worker := createWorker(0)
	after := time.After(time.Second * 10)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tick:
			fmt.Println("queue len = ", len(values))
		case <-after:
			fmt.Println("bye")
			return
		}
	}

}
