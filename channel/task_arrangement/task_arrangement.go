package main

import "fmt"

// 通过go routine 和 channel实现任务编排

/**
  编程题：
  3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，
	按照cat、dog、fish顺序打印在屏幕上100次。
*/
func control() {

	catChan := make(chan struct{})
	dogChan := make(chan struct{})
	fishChan := make(chan struct{})
	done := make(chan string)

	printCat(catChan, dogChan)
	printDog(dogChan, fishChan)
	printFish(fishChan, catChan, done)

	catChan <- struct{}{}

	fmt.Println(<-done)

}

func printCat(catChan, dogChan chan struct{}) {
	go func() {
		for i := 0; i < 100; i++ {
			if _, ok := <-catChan; ok {
				fmt.Println("cat")
				dogChan <- struct{}{}
			}
		}
	}()
}

func printDog(dogChan, fishChan chan struct{}) {
	go func() {
		for i := 0; i < 100; i++ {
			if _, ok := <-dogChan; ok {
				fmt.Println("dog")
				fishChan <- struct{}{}
			}
		}
	}()
}

func printFish(fishChan, catChan chan struct{}, done chan string) {
	go func() {
		for i := 0; i < 100; i++ {
			if _, ok := <-fishChan; ok {
				fmt.Println("fish")

				if i == 99 {
					// handle last task.
					done <- "all tasks done!"
				} else {
					catChan <- struct{}{}
				}

			}
		}

	}()
}

func main() {

	control()

}
