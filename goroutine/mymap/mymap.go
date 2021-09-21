package main

import "sync"

type MyMap struct {
	m map[string]int
}

//func NewMyMap() *MyMap {
//	return &MyMap{
//		m : make(map[string]int, 10),
//	}
//}
//
//func (m *MyMap) Put(key string, value interface{}) {
//	m.m[key] = value
//}
//
//func (m *MyMap) Get(key string) interface{} {
//	return m.Get(key)
//}

func main() {
	var m = sync.Map{}
	go func() {
		for {
			//m[1] = 1 //设置key
			m.Store(1, 1)
		}
	}()

	go func() {
		for {
			//_ = m[2] //访问这个map
			m.Load(2)
		}
	}()
	select {}
}
