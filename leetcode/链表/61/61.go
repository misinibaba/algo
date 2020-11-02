package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	dummy := ListNode{0, head}
	fast := &dummy
	slow := &dummy

	listLen := 0
	for fast.Next != nil {
		fast = fast.Next
		listLen++
	}

	n := k % listLen
	for i := 0; i < listLen - n; i++ {
		slow = slow.Next
	}
	fast.Next = head
	head = slow.Next
	slow.Next = nil
	return head
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{1,2}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := rotateRight(&l1, 1)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
