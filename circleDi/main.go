package main

import (
	"learngo/circleDi/A"
	"learngo/circleDi/B"
)

//	import cycle not allowed
//	package learngo/circleDi
//	imports learngo/circleDi/A
//	imports learngo/circleDi/B
//	imports learngo/circleDi/A
func main() {
	//A.Foo("good")

	a := &A.A{}
	b := &B.B{}
	a.Set(b)
	b.Set(a)

	a.Foo("goood")
}
