/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 *
 * https://leetcode-cn.com/problems/invert-binary-tree/description/
 *
 * algorithms
 * Easy (78.79%)
 * Likes:    1053
 * Dislikes: 0
 * Total Accepted:    322.5K
 * Total Submissions: 409.1K
 * Testcase Example:  '[4,2,7,1,3,6,9]'
 *
 * 翻转一棵二叉树。
 *
 * 示例：
 *
 * 输入：
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 2     7
 * ⁠/ \   / \
 * 1   3 6   9
 *
 * 输出：
 *
 * ⁠    4
 * ⁠  /   \
 * ⁠ 7     2
 * ⁠/ \   / \
 * 9   6 3   1
 *
 * 备注:
 * 这个问题是受到 Max Howell 的 原问题 启发的 ：
 *
 * 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func invertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	root.Left, root.Right = root.Right, root.Left
// 	invertTree(root.Left)
// 	invertTree(root.Right)
// 	return root
// }
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}

// @lc code=end

