package main

import (
	"fmt"
	"learngo/tree"
)

// 拓展已有结构体Node
// (组合的方式)拓展实现后序遍历
// 首字母小写，私有的
type myTreeNode struct {
	node *tree.Node
}

// 内嵌的方式拓展已有类型Node,
// 相当于把Node里面的变量、方法都拉到myTreeNodEmbedded里来。
type myTreeNodEmbedded struct {
	*tree.Node // embedding 内嵌的方式
}

func (myTreeNode *myTreeNodEmbedded) Traverse() {
	// 类似重载
	fmt.Println("This is a Traverse method!")
}

func (myNode *myTreeNode) postOrder() {

	if myNode == nil || myNode.node == nil {
		return
	}

	myNode.node.Left.Traverse()
	myNode.node.Right.Traverse()
	myNode.node.Print()

}

func main() {
	// 声明root变量为treeNode类型
	var root tree.Node

	// 初始化treeNode为value=3 left riht 为nil
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Traverse()

	count := 0
	root.TraverseFunc(func(n *tree.Node) {
		count++
	})
	fmt.Println("Count: ", count)

	// 用自己拓展的myTreeNode 做后序遍历
	//myNode := myTreeNode{&root}
	//myNode.postOrder()

	myNode := myTreeNodEmbedded{&root}
	// 调用自己重写的方法
	myNode.Traverse()
	// 调用原来Node 的Traverse方法
	myNode.Node.Traverse()
}
