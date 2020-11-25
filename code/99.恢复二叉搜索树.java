/*
 * @lc app=leetcode.cn id=99 lang=java
 *
 * [99] 恢复二叉搜索树
 *
 * https://leetcode-cn.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (62.01%)
 * Likes:    372
 * Dislikes: 0
 * Total Accepted:    43K
 * Total Submissions: 69.2K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * 给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
 * 
 * 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [1,3,null,null,2]
 * 输出：[3,1,null,null,2]
 * 解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [3,1,4,null,null,2]
 * 输出：[2,1,4,null,null,3]
 * 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树上节点的数目在范围 [2, 1000] 内
 * -2^31 
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
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    private void swap(TreeNode x, TreeNode y) {
        int tmp = x.val;
        x.val = y.val;
        y.val = tmp;
    }
    public void recoverTree(TreeNode root) {
        // for swap purpose
        TreeNode x = null, y = null;
        
        TreeNode predecessor = null, prev = null;
        while (root != null) {
            if (root.left != null) {
                // find predecessor
                predecessor = root.left;
                while (predecessor.right != null && predecessor.right != root) {
                    predecessor = predecessor.right;
                }
                // if no right child, then point to root
                if (predecessor.right == null) {
                    predecessor.right = root;
                    root = root.left;
                } else {
                    // has right child, it already point to root.
                    if (prev != null && root.val <= prev.val) {
                        y = root;
                        if (x == null) {
                            x = prev;
                        }
                    }
                    // store the root as prev
                    prev = root;
                    // recover from linking
                    predecessor.right = null;
                    // iterate right child
                    root = root.right;
                }
                continue;
            }

            // root.left is null
            if (prev != null && root.val <= prev.val) {
                y = root;
                if (x == null) {
                    x = prev;
                }
            }
            prev = root;
            root = root.right;
        }
        swap(x, y);
    }
}
// @lc code=end

