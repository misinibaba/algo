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
	pre := math.MinInt32
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		dfs(node.Left)
		if node.Val < pre {
			return false
		}
		pre = node.Val
		dfs(node.Right)

		return true
	}
	return dfs(root)
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
