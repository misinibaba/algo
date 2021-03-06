package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy.Next
	second := dummy.Next

	for n > 0 {
		second = second.Next
		n--
	}

	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return head
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{1,2,3,4,5}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := removeNthFromEnd(&l1, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
