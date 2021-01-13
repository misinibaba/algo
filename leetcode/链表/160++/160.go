package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA,curB := headA,headB
	for curA != curB {
		if curA == nil {    // 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil == nil
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA

}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{4, 1, 8, 4, 5}
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
	nums = []int{5, 0, 1, 8, 4, 5}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l2 = *l2.Next

	res := getIntersectionNode(&l1, &l2)
	fmt.Println(res)
}
