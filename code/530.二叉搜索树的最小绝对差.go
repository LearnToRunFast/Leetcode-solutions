/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 *
 * https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/description/
 *
 * algorithms
 * Easy (62.16%)
 * Likes:    286
 * Dislikes: 0
 * Total Accepted:    82.2K
 * Total Submissions: 132.3K
 * Testcase Example:  '[4,2,6,1,3]'
 *
 * 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
 *
 * 差值是一个正数，其数值等于两值之差的绝对值。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [4,2,6,1,3]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,0,48,null,null,12,49]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点的数目范围是 [2, 10^4]
 * 0 <= Node.val <= 10^5
 *
 *
 *
 *
 * 注意：本题与 783
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
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
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// func getMinimumDifference(root *TreeNode) int {
// 	if root == nil {
// 		return -1
// 	}
// 	ans := 1<<31 - 1
// 	var prev *TreeNode
// 	stack := []*TreeNode{root}
// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]
// 		if node != nil {
// 			if node.Right != nil {
// 				stack = append(stack, node.Right)
// 			}
// 			stack = append(stack, node, nil)
// 			if node.Left != nil {
// 				stack = append(stack, node.Left)
// 			}
// 		} else {
// 			node = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 			if prev != nil {
// 				tmp := abs(node.Val - prev.Val)
// 				ans = min(ans, tmp)
// 			}
// 			prev = node
// 		}
// 	}
// 	return ans
// }

// recursion
var prev *TreeNode
var ans int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if prev != nil {
		tmp := abs(root.Val - prev.Val)
		ans = min(ans, tmp)

	}
	prev = root
	dfs(root.Right)
}
func getMinimumDifference(root *TreeNode) int {
	ans = 1<<31 - 1
	dfs(root)
	return ans
}

// @lc code=end

