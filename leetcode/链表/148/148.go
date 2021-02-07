package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{4,2,1,3}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := sortList(&l1)
	fmt.Println(res)
}
