package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)


func variableZeroValue()  {

	var a int
	var s string
	fmt.Printf("%d %q\n", a ,s )
	
}

func variableInitailValue(){

	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a , b ,s)
}

func variableTypeDeduction()  {

	var a,b,c, s = 3,4,true,"def"
	fmt.Println(a , b ,c, s)
}

func variableShorter()  {

	a,b,c, s := 3,4,true,"def"
	b = 5
	fmt.Println(a , b ,c, s)
}

//验证欧拉公式
func euler(){
	//声明一个变量c为复数
	c := 3 + 4i
	//输出取绝对值 |3 + 4i| = 5
	fmt.Println(cmplx.Abs(c))
}

func triangle(){
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const filename string = "abc.txt"
	const a,b  = 3,4
	var c int
	//这里不需要用float64做强制转换，因为常量相当于文本替换的动作，a,b没有
	//声明类型，可以做int 也可以做 float
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename , c)

}

func enum(){

	const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)

	//自增值枚举类型
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}


func main() {

	fmt.Println("hello")
	variableZeroValue()
	variableInitailValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enum()
}
