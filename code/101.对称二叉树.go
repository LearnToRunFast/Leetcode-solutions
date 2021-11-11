/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (56.50%)
 * Likes:    1600
 * Dislikes: 0
 * Total Accepted:    435.9K
 * Total Submissions: 771.3K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 *
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
// func isSymmetricCore(left *TreeNode, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	} else if left == nil || right == nil {
// 		return false
// 	} else if left.Val != right.Val {
// 		return false
// 	}

// 	return isSymmetricCore(left.Left, right.Right) && isSymmetricCore(left.Right, right.Left)
// }

// func isSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}
// 	return isSymmetricCore(root.Left, root.Right)
// }

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}
	return true
}

// @lc code=end

