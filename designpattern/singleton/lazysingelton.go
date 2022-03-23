package singleton

import "sync"

/**
	懒汉模式  并发安全，双重判断
 */
type lazySingleton struct {

}

var lazyIns *lazySingleton

var mutex *sync.Mutex

func GetLazySingleton() *lazySingleton{

	if lazyIns == nil{
		mutex.Lock()
		if lazyIns == nil{
			lazyIns = &lazySingleton{}
		}
		mutex.Unlock()
	}

	return lazyIns
}
