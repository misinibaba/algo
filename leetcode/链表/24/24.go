package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	first := head
	second := head.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{1,2,3,4}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := swapPairs(&l1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
