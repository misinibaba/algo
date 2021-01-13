package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	var ans []int
	dfs = func(node *TreeNode) {
		if len(ans) == k || node == nil {
			return
		}

		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ans[k - 1]
}

func main() {
	var t1 TreeNode
	var genTree func(int, *TreeNode) *TreeNode
	nums := []int{3, 1, 4, -1 ,2}
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

	res := kthSmallest(&t1, 1)
	fmt.Println(res)
}
