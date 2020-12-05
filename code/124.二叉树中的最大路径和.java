/*
 * @lc app=leetcode.cn id=124 lang=java
 *
 * [124] 二叉树中的最大路径和
 *
 * https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/description/
 *
 * algorithms
 * Hard (43.19%)
 * Likes:    795
 * Dislikes: 0
 * Total Accepted:    86K
 * Total Submissions: 198.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空二叉树，返回其最大路径和。
 * 
 * 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：[1,2,3]
 * 
 * ⁠      1
 * ⁠     / \
 * ⁠    2   3
 * 
 * 输出：6
 * 
 * 
 * 示例 2：
 * 
 * 输入：[-10,9,20,null,null,15,7]
 * 
 * -10
 * / \
 * 9  20
 * /  \
 * 15   7
 * 
 * 输出：42
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode() {} TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) { this.val = val; this.left
 * = left; this.right = right; } }
 */
class Solution {
    int max = Integer.MIN_VALUE;

    private int dfs(TreeNode root) {
        if (root == null)
            return 0;

        int left = Math.max(0, dfs(root.left));
        int right = Math.max(0, dfs(root.right));
        int weight = root.val + left + right;
        max = Math.max(weight, max);

        return root.val + Math.max(left, right);
    }

    public int maxPathSum(TreeNode root) {
        dfs(root);
        return max;
    }
}
// @lc code=end
