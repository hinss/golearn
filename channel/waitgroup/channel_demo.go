package main

import (
	"fmt"
)

// 使用done channel 等待任务结束

func doWorker(id int, w worker) {
	for {
		fmt.Printf("Workder:%d received %c\n",
			id, <-w.c)

		go func() {
			w.done <- true
		}()

	}
}

// channel 作为一等公民也可以作为函数的返回值、参数传递
func createWorker(i int) worker {
	w := worker{
		c:    make(chan int),
		done: make(chan bool),
	}

	go doWorker(i, w)
	return w
}

type worker struct {
	c    chan int
	done chan bool
}

func chanDemo() {

	// channel队列，创建10个只能用于发送数据的channel。
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, w := range workers {
		w.c <- 'a' + i
	}

	for _, w := range workers {
		<-w.done
	}

	for i, w := range workers {
		w.c <- 'A' + i
	}

	for _, w := range workers {
		<-w.done
	}

}

func main() {
	chanDemo()

}
