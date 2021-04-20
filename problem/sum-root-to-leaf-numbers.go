package problem

/*
问题：
129. 求根节点到叶节点数字之和
https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
*/

func sumNumbers(root *TreeNode) int {
	return sumNumbersWithPrefix(root, 0)
}

func sumNumbersWithPrefix(root *TreeNode, pre int) int {
	if root == nil {
		return 0
	}
	now := pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return now
	}
	return sumNumbersWithPrefix(root.Left, now) + sumNumbersWithPrefix(root.Right, now)
}
