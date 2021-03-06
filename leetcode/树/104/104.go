package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level + 1 > ans {
			ans = level + 1
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
	nums := []int{3,9,20,-1,-1,15,7}
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

	res := maxDepth(&t1)
	fmt.Println(res)
}
