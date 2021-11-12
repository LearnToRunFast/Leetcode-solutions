/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (70.51%)
 * Likes:    1281
 * Dislikes: 0
 * Total Accepted:    261.1K
 * Total Submissions: 370.3K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
 *
 *
 *
 * 示例 1:
 *
 *
 * Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
 * Output: [3,9,20,null,null,15,7]
 *
 *
 * 示例 2:
 *
 *
 * Input: preorder = [-1], inorder = [-1]
 * Output: [-1]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * inorder.length == preorder.length
 * -3000
 * preorder 和 inorder 均无重复元素
 * inorder 均出现在 preorder
 * preorder 保证为二叉树的前序遍历序列
 * inorder 保证为二叉树的中序遍历序列
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 1. if it's empty, return nil
	if len(preorder) == 0 {
		return nil
	}
	// 2. get the root value, it will be the first element of preorder
	root := &TreeNode{Val: preorder[0]}

	// 3. find the root in inorder
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
			break
		}
	}

	// 4. build left and right subtree
	// the length of preorder and inorder should be the same
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}

// @lc code=end

