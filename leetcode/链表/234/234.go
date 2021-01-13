package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast := head
	second := head
	for fast != nil && fast.Val != 0 && fast.Next != nil && fast.Next.Val != 0{
		fast = fast.Next
		second = second.Next.Next
	}

	secondList := reverse(fast)


	for secondList != nil && secondList.Val != 0 {
		if head.Val != secondList.Val {
			return false
		}
		head = head.Next
		secondList = secondList.Next
	}

	return true
}
func reverse(head *ListNode) *ListNode {
	prev := new(ListNode)
	cur := head
	for cur != nil && cur.Val != 0{
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{-129, -129}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := isPalindrome(&l1)
	fmt.Println(res)
}
