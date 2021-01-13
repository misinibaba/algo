package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var dfs func(*TreeNode)
	var ans []int
	dfs = func(node *TreeNode) {
		if node == nil || node.Val == 0 {
			return
		}

		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}
	dfs(root)
	return ans
}


func main() {
	var t1 TreeNode
	var genTree func(int, *TreeNode) *TreeNode
	nums := []int{1, -1, 2, 3}
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

	res := preorderTraversal(&t1)
	fmt.Println(res)
}
