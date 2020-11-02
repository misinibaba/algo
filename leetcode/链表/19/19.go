package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// dummy用于在只有一个值的时候的截取
	dummy := new(ListNode)
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 0; i <= n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	node := new(ListNode)
	var l1 ListNode
	l1.Next = node
	nums := []int{1}
	for k, v := range nums {
		node.Val = v
		if k < len(nums) - 1 {
			node.Next = new(ListNode)
			node = node.Next
		}
	}
	l1 = *l1.Next

	res := removeNthFromEnd(&l1, 1)
	for res != nil {
		//fmt.Println(res.Val)
		res = res.Next
	}
}
