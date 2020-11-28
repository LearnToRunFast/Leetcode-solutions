import java.util.HashMap;
import java.util.LinkedList;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=105 lang=java
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inOrderLeft-traversal/description/
 *
 * algorithms
 * Medium (68.50%)
 * Likes:    770
 * Dislikes: 0
 * Total Accepted:    131.6K
 * Total Submissions: 192K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 * 
 * 注意:
 * 你可以假设树中没有重复的元素。
 * 
 * 例如，给出
 * 
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inOrderLeft = [9,3,15,20,7]
 * 
 * 返回如下的二叉树：
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode(int x) { val = x; } }
 */
// class Solution {
// private TreeNode buildTree(int[] preorder, int[] inorder, int preOrderLeft,
// int preOrderRight, int inOrderLeft,
// int inOrderRight, Map<Integer, Integer> lookup) {
// if (preOrderLeft > preOrderRight) {
// return null;
// }
// TreeNode root = new TreeNode(preorder[preOrderLeft]);
// int inOrderRoot = lookup.get(preorder[preOrderLeft]);
// int sizeOfLeftTree = inOrderRoot - inOrderLeft;

// root.left = buildTree(preorder, inorder, preOrderLeft + 1, preOrderLeft +
// sizeOfLeftTree, inOrderLeft,
// inOrderRoot - 1, lookup);
// root.right = buildTree(preorder, inorder, preOrderLeft + sizeOfLeftTree + 1,
// preOrderRight, inOrderRoot + 1,
// inOrderRight, lookup);
// return root;
// }

// public TreeNode buildTree(int[] preorder, int[] inorder) {
// Map<Integer, Integer> lookup = new HashMap<>();
// for (int i = 0; i < inorder.length; i++) {
// lookup.put(inorder[i], i);
// }
// return buildTree(preorder, inorder, 0, preorder.length - 1, 0, inorder.length
// - 1, lookup);
// }
// }
class Solution {

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder == null || preorder.length == 0)
            return null;

        TreeNode root = new TreeNode(preorder[0]);
        LinkedList<TreeNode> stack = new LinkedList<>();
        stack.push(root);
        int inOrderCurr = 0;
        for (int preOrderCurr = 1; preOrderCurr < preorder.length; preOrderCurr++) {
            TreeNode node = stack.peek();
            if (node.val != inorder[inOrderCurr]) {
                node.left = new TreeNode(preorder[preOrderCurr]);
                stack.push(node.left);
            } else {
                while (!stack.isEmpty() && stack.peek().val == inorder[inOrderCurr]) {
                    node = stack.pop();
                    inOrderCurr++;
                }
                node.right = new TreeNode(preorder[preOrderCurr]);
                stack.push(node.right);
            }

        }
        return root;
    }
}
// @lc code=end
