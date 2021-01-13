package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	for (root != nil && root.Val != 0) || len(stack) > 0 {
		for root != nil && root.Val != 0 {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return
}

func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || node.Val == 0 {
			return
		}

		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

func main() {
	var t1 TreeNode
	var genTree func(int, *TreeNode) *TreeNode
	nums := []int{1, -1, 2, -1, -1, 3}
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
	res := inorderTraversal(&t1)
	fmt.Println(res)
	res = inorderTraversal2(&t1)
	fmt.Println(res)
}
