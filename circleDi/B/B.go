package B

import (
	"learngo/circleDi/service"
)

type B struct {
	a service.A
}

func (b *B) Set(a service.A) {
	b.a = a
}

func (b *B) Goo(a string) (string) {
	return b.a.Minus(a)
}

func (b *B) Add(a string) string {
	return a + "----"
}
