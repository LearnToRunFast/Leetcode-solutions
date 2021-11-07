/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (70.38%)
 * Likes:    674
 * Dislikes: 0
 * Total Accepted:    434.9K
 * Total Submissions: 617.5K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,null,2,3]
 * 输出：[1,2,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = [1]
 * 输出：[1]
 *
 *
 * 示例 4：
 *
 *
 * 输入：root = [1,2]
 * 输出：[1,2]
 *
 *
 * 示例 5：
 *
 *
 * 输入：root = [1,null,2]
 * 输出：[1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数目在范围 [0, 100] 内
 * -100
 *
 *
 *
 *
 * 进阶：递归算法很简单，你可以通过迭代算法完成吗？
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

// recursion
// func recursion(root *TreeNode, ans *[]int) {
// 	if root == nil {
// 		return
// 	}
// 	*ans = append(*ans, root.Val)
// 	recursion(root.Left, ans)
// 	recursion(root.Right, ans)
// }
// func preorderTraversal(root *TreeNode) []int {
// 	ans := make([]int, 0)
// 	recursion(root, &ans)
// 	return ans
// }

// stack version
// func preorderTraversal(root *TreeNode) []int {
// 	ans := make([]int, 0)
// 	stack := make([]*TreeNode, 0)
// 	if root != nil {
// 		stack = append(stack, root)
// 	}
// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		ans = append(ans, node.Val)
// 		if node.Right != nil {
// 			stack = append(stack, node.Right)
// 		}
// 		if node.Left != nil {
// 			stack = append(stack, node.Left)
// 		}
// 	}
// 	return ans
// }
// general form
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	ans := make([]int, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
		}
	}
	return ans
}

// @lc code=end

