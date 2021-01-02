package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printFile(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	//没有赋值，没有递增条件，就相当于while循环的效果
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

//死循环
func forever() {

	for {
		fmt.Println("abc")
	}
}

//省略初始条件，相当于while
func convertToBin(v int) {

	result := ""
	for ; v > 0; v /= 2 {
		result = strconv.Itoa(v%2) + result
	}

	fmt.Println(result)

}

func main() {

	printFile("abc.txt")

	//forever()

	convertToBin(5)
}
