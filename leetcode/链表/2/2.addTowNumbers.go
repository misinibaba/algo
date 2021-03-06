package main

import "fmt"

type ListNode struct {
 	Val int
 	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans = new(ListNode)
	head := ans
	var sum, fix int
	for l1 != nil || l2 != nil{
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum = v1 + v2 + fix
		ans.Next = &ListNode{sum % 10, nil}
		ans = ans.Next
		fix = sum / 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	if fix > 0 {
		ans.Next = &ListNode{fix, nil}
	}
	return head.Next
}

func main() {
	preNode := new(ListNode)
	var l1 ListNode
	l1.Next = preNode
	for _, v := range []int{9,9,9,9,9,9,9} {
		preNode.Val = v
		preNode.Next = new(ListNode)
		preNode = preNode.Next
	}
	l1 = *l1.Next

	preNode = new(ListNode)
	var l2 ListNode
	l2.Next = preNode
	for _, v := range []int{9,9,9,9} {
		preNode.Val = v
		preNode.Next = new(ListNode)
		preNode = preNode.Next
	}
	l2 = *l2.Next

	res := addTwoNumbers(&l1, &l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
