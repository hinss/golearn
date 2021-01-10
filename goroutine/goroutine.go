package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	var a [10]int
	for i := 0; i < 10; i++ {
		// 匿名函数 go 关键字开启协程
		go func(i int) { // race condition
			for {
				// 无限循环的进行IO 操作，所以它会交出cpu的控制权(主动的)
				//fmt.Printf("Hello from gorounine %d\n", i)

				// 而这种非IO的操作，是没有进行协程之间的切换的，所以并不一定每个协程都有执行的机会
				a[i]++
				// 主动让出执行权，一般不需要这样
				// 加了之后发现数组中的数值反而自增的少了许多
				runtime.Gosched()
			}
		}(i)
	}
	// 不加这句 main结束就会杀掉所有的协程 所以什么都不会打
	// 加这句表示在main睡1ms，在这ims内看协程打印了多少东西
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
