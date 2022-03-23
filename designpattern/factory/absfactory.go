package factory

import "fmt"

/**
抽象工厂  返回的是接口而非结构体
简单工厂就是简单的返回类
可以通过拓展实现多个工厂函数 来返回不同的接口实现
*/

type IPerson interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

func NewIPerson(name string, age int) IPerson {
	return person{
		name: name,
		age: age,
	}
}
