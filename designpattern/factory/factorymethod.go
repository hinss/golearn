package factory

/**
工厂方法模式 返回的是创建对象的函数
 */

type PersonNew struct {
	Name string
	Age int
}

type PersonNewFunc func(name string) PersonNew

func NewPersonNewFactory(age int) PersonNewFunc  {
	return func(name string) PersonNew {
		return PersonNew{
			Name: name,
			Age: age,
		}
	}
}