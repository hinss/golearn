package main

import (
	"context"
	"fmt"
	"time"
)

/**
 	context的用法
 */

func RedisProcess(ctx context.Context) {

	for {
		select {
		case r := <-ctx.Done():
			fmt.Println("Receive from channel:", r)
			fmt.Println("Redis process done...")
			return
		default:
			fmt.Println("Redis is processing....")
			time.Sleep(time.Second)
		}
	}
}

func MysqlProcess(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Mysql process done...")
			return
		default:
			fmt.Println("Mysql is processing....")
			time.Sleep(time.Second)
		}
	}

}



func HandleRequest(ctx context.Context) {

	fmt.Println("Handle request....")

	go RedisProcess(ctx)
	go MysqlProcess(ctx)

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)

	time.Sleep(time.Second * 5)
	// 调用取消方法传递关闭所有go routine
	cancel()

	time.Sleep(time.Second * 10)

}
