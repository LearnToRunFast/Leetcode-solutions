/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (51.47%)
 * Likes:    375
 * Dislikes: 0
 * Total Accepted:    68.9K
 * Total Submissions: 133.9K
 * Testcase Example:  '[1,null,2,2]'
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 *
 *
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 * 返回[2].
 *
 * 提示：如果众数超过1个，不需考虑输出顺序
 *
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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
var res []int
var maxCount int
var pre *TreeNode
var count int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Left)
	if pre != nil && root.Val == pre.Val {
		count++
	} else {
		count = 1
	}
	pre = root
	if count > maxCount {
		maxCount = count
		res = []int{root.Val}
	} else if count == maxCount {
		res = append(res, root.Val)
	}
	dfs(root.Right)
}
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	maxCount = 0
	count = 0
	dfs(root)
	return res
}

// @lc code=end

