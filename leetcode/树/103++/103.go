package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var dfs func(node *TreeNode, level int)
	ans := make([][]int, 0)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(ans) == level {
			ans = append(ans, []int{})
		}

		if level % 2 == 0 {
			ans[level] = append(ans[level], node.Val)
		} else {
			ans[level] = append([]int{node.Val}, ans[level]...)
		}
		dfs(node.Left, level + 1)
		dfs(node.Right, level + 1)
	}
	dfs(root, 0)
	return ans
}


func main() {
	var t1 TreeNode
	var genTree func(int, *TreeNode) *TreeNode
	nums := []int{3, 9, 20, -1, -1, 15, 7}
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

	res := zigzagLevelOrder(&t1)
	fmt.Println(res)
}
