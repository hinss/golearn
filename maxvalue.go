package main

import "fmt"

func max(vals...int) int {

	var max int

	for _, val := range vals {
		fmt.Print(val)
	}

	return max

}





func main() {

	fmt.Println(max(1,2,3,99,4))
}
