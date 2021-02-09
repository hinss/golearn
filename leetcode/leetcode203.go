package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	// 构造虚拟头节点
	dummyHead := &ListNode{
		Val: 0,
	}

	cur := dummyHead
	cur.Next = head

	for head != nil {

		if head.Val == val {
			cur.Next = head.Next
			head = head.Next
		} else {
			head = head.Next
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

func main() {

	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}

	head = removeElements(head, 6)

}
