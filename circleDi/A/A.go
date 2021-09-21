package A

import (
	"learngo/circleDi/service"
	"strings"
)

type A struct {
	b service.B
}

func (a *A) Set(b service.B) {
	a.b = b
}

func (a *A) Minus(s string) string {
	return strings.Trim(s, "\t")
}

func (a *A) Foo(s string) (string) {
	return a.b.Add(s)
}
