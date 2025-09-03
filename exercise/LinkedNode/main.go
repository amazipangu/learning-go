package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middleHead := head
	count := 0

	for head != nil {
		count++
		head = head.Next
	}
	middleIndex := int(count / 2)
	for i := 0; i < middleIndex; i++ {
		middleHead = middleHead.Next
	}

	return middleHead
}

func main() {
	node3 := &ListNode{Val: 3, Next: nil}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}
	middleNode(node1)
}
