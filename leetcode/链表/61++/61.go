package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

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

	res := rotateRight(&l1, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
