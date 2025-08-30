package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if -100 <= root.Val && root.Val <= 100 {
		var total int = root.Left.Val + root.Right.Val
		if root.Val == total {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
