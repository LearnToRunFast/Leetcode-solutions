/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 *
 * https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (69.00%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    77.4K
 * Total Submissions: 112.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 输出：[3, 14.5, 11]
 * 解释：
 * 第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 节点值的范围在32位有符号整数范围内。
 *
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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var res []float64
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var sum float64
		size := len(queue)
		for i := 0; i < size; i++ {
			sum += float64(queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, sum/float64(size))
		queue = queue[size:]
	}
	return res
}

// @lc code=end

