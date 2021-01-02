package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 给treeNode结构体定义一些方法
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 由于go语言都是值传递 所有这里要接收treeNode的指针
func (node *Node) SetValue(value int) {
	node.Value = value
}

// 使用自定义工厂函数
func CreateNode(value int) *Node {
	// 这里返回局部变量的地址在c++会报错
	return &Node{Value: value}
}
