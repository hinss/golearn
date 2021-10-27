package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
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

func chanDemo() {

	// channel队列，创建10个只能用于发送数据的channel。
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
