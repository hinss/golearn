package commucate

import "fmt"

func chanDemoWithDescription() {
	// channel的定义 channel也是一等公民 可以是参数也能做函数返回值
	// 定义个叫做c channel 可以装int类型的数据 此时为nil
	//var c chan int
	// 一般用make来创建channel
	c := make(chan int)
	// 往channel发数据 这里会挂掉  channel必须是channel之间的操作 也就是要有另外一个channel去收这个丢到c的1、2...
	c <- 1
	c <- 2
	// 收channel发送的数据
	n := <-c
	fmt.Println(n)
}

func chanDemo() {

	c := make(chan int)
	go func() {
		// 这里就是定义一个收c的地方
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
}
