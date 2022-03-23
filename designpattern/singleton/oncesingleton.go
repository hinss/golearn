package singleton

import (
	"fmt"
	"sync"
)

type onceSingleton struct {

}

var onceIns *onceSingleton

var once sync.Once

func GetOnceSingleton() *onceSingleton{

	once.Do(func() {
		fmt.Println("init onceSingleton instance...")
		onceIns = &onceSingleton{}
	})

	return onceIns
}

func GetOnceSingleton2() *onceSingleton{

		fmt.Println("init onceSingleton instance...")
		onceIns = &onceSingleton{}

	return onceIns
}
