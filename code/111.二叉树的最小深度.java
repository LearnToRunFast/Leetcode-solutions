import java.util.Deque;
import java.util.LinkedList;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=111 lang=java
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (45.16%)
 * Likes:    410
 * Dislikes: 0
 * Total Accepted:    158.4K
 * Total Submissions: 348.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明：叶子节点是指没有子节点的节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000 
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode() {} TreeNode(int val) { this.val = val; }
 * TreeNode(int val, TreeNode left, TreeNode right) { this.val = val; this.left
 * = left; this.right = right; } }
 */
// class Solution {
// public int minDepth(TreeNode root) {
// if (root == null)
// return 0;
// int ans = 0;
// Deque<TreeNode> queue = new LinkedList<>();
// queue.add(root);
// while (!queue.isEmpty()) {
// int size = queue.size();
// ans++;
// for (int i = 0; i < size; i++) {
// TreeNode node = queue.poll();
// if (node.left == null && node.right == null) {
// return ans;
// }
// if (node.left != null) {
// queue.add(node.left);
// }

// if (node.right != null) {
// queue.add(node.right);
// }
// }
// }
// return ans;
// }
// }
class Solution {
    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int left = minDepth(root.left);
        int right = minDepth(root.right);

        return root.left == null || root.right == null ? left + right + 1 : Math.min(left, right) + 1;
    }
}
// @lc code=end
