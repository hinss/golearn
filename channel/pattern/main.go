package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
				i++
			case <-done:
				// 模拟优雅退出
				fmt.Println("clean up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleanup done")
				return
			}

		}
	}()
	return c
}

func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

// 非阻塞等待
func noneBlockingWait(c chan string) (string, bool) {

	select {
	case msg := <-c:
		return msg, true
	default:
		return "", false
	}
}

// 超时控制
func timeOutWait(c chan string, timeout time.Duration) (string, bool) {

	select {
	case msg := <-c:
		return msg, true
	case <-time.After(timeout):
		return "", false
	}
}

func main() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	//m2 := msgGen("service2", done)
	//m3 := msgGen("service3")
	//m := fanIn(m1, m2)
	//for {
	//	fmt.Println(<-m)
	//}
	for i := 0; i < 5; i++ {
		if m, ok := timeOutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}

	// 优雅退出，可以利用channel双向的特性，给done送数据，通知任务要退出
	done <- struct{}{}
	// 任务退出后往done送数据，主任务等待done送回来的数据再退出
	<-done

}
