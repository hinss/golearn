package main

import (
	"fmt"
	"runtime"
)

func main() {

	for i := 0; i<4 ; i++ {
		getResult()

		fmt.Println("goroutines: ", runtime.NumGoroutine())

	}
	
}

func getResult() int {

	ch := make(chan int)
	for i:=0; i<3; i++ {
		go func() {
			ch<- i
		}()
	}
	return <-ch
}
