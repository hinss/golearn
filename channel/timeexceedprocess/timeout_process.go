package main

import (
	"context"
	"fmt"
	"time"
)

// 练习使用  go routine + context + select 完成 go routine 超时处理
func main() {
	withTimeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	go handleRequest(withTimeout, 500*time.Millisecond)

	select {
	case <-withTimeout.Done():
		fmt.Println("timeout in main function!")
	}
}

func handleRequest(ctx context.Context, duration time.Duration) {

	select {
	case <-ctx.Done():
		fmt.Println("request timeout", ctx.Err())
	case <-time.After(duration):
		fmt.Println("finished request")
	}
}
