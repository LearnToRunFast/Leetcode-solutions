/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (72.17%)
 * Likes:    615
 * Dislikes: 0
 * Total Accepted:    140.7K
 * Total Submissions: 195.1K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 1. if it's empty, return nil
	if len(inorder) == 0 {
		return nil
	}
	// 2. if it's not empty, then the root will be the last element of postorder
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// 3. find the root in inorder
	rootIndex := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			rootIndex = i
			break
		}
	}

	// 4. build left and right subtree
	// inorder tree  and postorder tree must be the same length
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}

// @lc code=end

