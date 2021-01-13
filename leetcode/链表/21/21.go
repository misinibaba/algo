package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{1,2,4}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	node = new(ListNode)
	var l2 ListNode
	l2.Next = node
	nums = []int{1,3,4}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l2 = *l2.Next

	res := mergeTwoLists(&l1, &l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
