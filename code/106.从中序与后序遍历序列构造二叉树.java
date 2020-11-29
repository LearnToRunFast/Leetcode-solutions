import java.util.Deque;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Map;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=106 lang=java
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (70.74%)
 * Likes:    409
 * Dislikes: 0
 * Total Accepted:    78.7K
 * Total Submissions: 111.1K
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
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode(int x) { val = x; } }
 */
// class Solution {
// int postIdx = 0;

// private TreeNode buildTree(int[] inorder, int[] postorder, int inorderLeft,
// int inorderRight,
// Map<Integer, Integer> lookup) {
// if (inorderLeft > inorderRight) {
// return null;
// }
// int rootVal = postorder[postIdx];
// TreeNode root = new TreeNode(rootVal);

// int inorderRootIdx = lookup.get(rootVal);
// postIdx--;
// root.right = buildTree(inorder, postorder, inorderRootIdx + 1, inorderRight,
// lookup);
// root.left = buildTree(inorder, postorder, inorderLeft, inorderRootIdx - 1,
// lookup);
// return root;
// }

// public TreeNode buildTree(int[] inorder, int[] postorder) {
// Map<Integer, Integer> lookup = new HashMap<>();
// postIdx = postorder.length - 1;
// for (int i = 0; i < inorder.length; i++) {
// lookup.put(inorder[i], i);
// }
// return buildTree(inorder, postorder, 0, inorder.length - 1, lookup);
// }
// }
class Solution {

    public TreeNode buildTree(int[] inorder, int[] postorder) {
        if (postorder == null || postorder.length == 0) {
            return null;
        }
        TreeNode root = new TreeNode(postorder[postorder.length - 1]);
        Deque<TreeNode> stack = new LinkedList<>();
        stack.push(root);
        int inorderIdx = inorder.length - 1;
        for (int postorderIdx = postorder.length - 2; postorderIdx >= 0; postorderIdx--) {
            int postorderVal = postorder[postorderIdx];
            TreeNode node = stack.peek();
            if (node.val != inorder[inorderIdx]) {
                node.right = new TreeNode(postorderVal);
                stack.push(node.right);
            } else {

                while (!stack.isEmpty() && stack.peek().val == inorder[inorderIdx]) {
                    node = stack.pop();
                    inorderIdx--;
                }
                // if they are no longer equals, we know current value is a left child
                node.left = new TreeNode(postorderVal);
                stack.push(node.left);
            }
        }
        return root;
    }
}
// @lc code=end
