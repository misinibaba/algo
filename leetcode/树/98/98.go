package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(*TreeNode, int, int) bool
	check = func(node *TreeNode, low int, high int) bool {
		if node == nil {
			return true
		}

		if node.Val <= low || node.Val >= high {
			return false
		}

		return check(node.Left, low, node.Val) && check(node.Right, node.Val, high)
	}
	return check(root, math.MinInt64, math.MaxInt64)
}

func main() {
	var t1 TreeNode
	var genTree func(int, *TreeNode) *TreeNode
	nums := []int{2, 1, 3}
	genTree = func(index int, t *TreeNode) *TreeNode {
		if index > len(nums) - 1 || nums[index] == -1 {
			return t
		}

		t.Val = nums[index]
		t.Left  = genTree(2 * index + 1, new(TreeNode))
		t.Right = genTree(2 * index + 2, new(TreeNode))
		return t
	}

	genTree(0, &t1)
	res := isValidBST(&t1)
	fmt.Println(res)
}
