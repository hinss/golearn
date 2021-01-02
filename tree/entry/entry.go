package main

import (
	"fmt"
	"learngo/tree"
)

// (组合的方式)拓展实现后序遍历
// 首字母小写，私有的
type myTreeNode struct {
	node *tree.Node
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

}
