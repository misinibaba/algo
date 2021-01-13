package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level == len(ans) {
			ans = append(ans, []int{})
		}

		ans[level] = append(ans[level], node.Val)
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

	res := levelOrder(&t1)
	fmt.Println(res)
}
