package singleton

import (
	"fmt"
	"testing"
)

func TestGetOnceSingleton(t *testing.T) {

	for i := 0; i < 100; i++ {
		fmt.Println("goroutine running..")
		go func() {
			GetOnceSingleton()
			//GetOnceSingleton2()
		}()
	}

}
