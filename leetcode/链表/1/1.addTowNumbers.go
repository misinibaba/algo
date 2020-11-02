package main

import "fmt"

type ListNode struct {
 	Val int
 	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	resList := dummy
	sum := 0
	fix := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum = v1 + v2 + fix
		fix = sum / 10
		resList.Next = new(ListNode)
		resList.Next.Val = sum % 10

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

		resList = resList.Next
	}
	if fix != 0 {
		resList.Next = new(ListNode)
		resList.Next.Val = fix
	}

	return dummy.Next
}

func main() {
	preNode := new(ListNode)
	var l1 ListNode
	l1.Next = preNode
	for _, v := range []int{1,8} {
		preNode.Val = v
		preNode.Next = new(ListNode)
		preNode = preNode.Next
	}
	l1 = *l1.Next

	preNode = new(ListNode)
	var l2 ListNode
	l2.Next = preNode
	for _, v := range []int{0} {
		preNode.Val = v
		preNode.Next = new(ListNode)
		preNode = preNode.Next
	}
	l2 = *l2.Next

	res := addTwoNumbers(&l1, &l2)

	for _, k := range []int{1,2,3} {
		fmt.Println(res.Val)
		res = res.Next
		k++
	}
}
