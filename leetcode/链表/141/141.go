package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil || slow == nil {
			return false
		}
	}
	return true
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{3,2,0,-4}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := hasCycle(&l1)
	fmt.Println(res)
}
