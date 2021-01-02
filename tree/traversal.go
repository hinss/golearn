package tree

// 同一个Node struc的方法体可以定义到不同的文件但是package必须相同
// 中序遍历
//func (node *Node) Traverse() {
//	if node == nil {
//		return
//	}
//
//	// 这里不需要判断node.Left 是否为空 跟java不一样
//	// 只需要node 上面判断了就好了
//	node.Left.Traverse()
//	node.Print()
//	node.Right.Traverse()
//}

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	// 这里不需要判断node.Left 是否为空 跟java不一样
	// 只需要node 上面判断了就好了
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
