/*
 * @lc app=leetcode.cn id=404 lang=java
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (56.30%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    63.7K
 * Total Submissions: 113.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 * 
 * 示例：
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 * 
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
// class Solution {
//     private int helper(TreeNode root, TreeNode parent) {
//         if (root == null) return 0;
//         if (root.left == null && root.right == null) {
//             if (parent.left == root) {
//                 return root.val;
//             }
//             return 0;
//         }
//         return helper(root.left, root) + helper(root.right, parent);
        
//     }
//     public int sumOfLeftLeaves(TreeNode root) {
//         return helper(root, root);
//     }
// }
class Solution {
    private int helper(TreeNode root) {
        int ans = 0;
        if (root.left != null) {
            if (root.left.left == null && root.left.right == null) {
                ans += root.left.val;
            } else {
                ans += helper(root.left);
            }
        }
        if (root.right != null) {
            if (root.right.left != null || root.right.right != null) {
                ans += helper(root.right);
            }
        }
        return ans;
    }
    public int sumOfLeftLeaves(TreeNode root) {
        return root == null ? 0 : helper(root);
    }
}
// @lc code=end

