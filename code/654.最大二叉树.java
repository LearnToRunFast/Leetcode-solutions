/*
 * @lc app=leetcode.cn id=654 lang=java
 *
 * [654] 最大二叉树
 *
 * https://leetcode-cn.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (81.39%)
 * Likes:    227
 * Dislikes: 0
 * Total Accepted:    28.2K
 * Total Submissions: 34.6K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
 * 
 * 
 * 二叉树的根是数组中的最大元素。
 * 左子树是通过数组中最大值左边部分构造出的最大二叉树。
 * 右子树是通过数组中最大值右边部分构造出的最大二叉树。
 * 
 * 
 * 通过给定的数组构建最大二叉树，并且输出这个树的根节点。
 * 
 * 
 * 
 * 示例 ：
 * 
 * 输入：[3,2,1,6,0,5]
 * 输出：返回下面这棵树的根节点：
 * 
 * ⁠     6
 * ⁠   /   \
 * ⁠  3     5
 * ⁠   \    / 
 * ⁠    2  0   
 * ⁠      \
 * ⁠       1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给定的数组的大小在 [1, 1000] 之间。
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
    private int findMaxIndex(int[] nums, int l, int r) {
        int maxIdx = l;
        for (int i = l + 1; i < r; i++) {
            if (nums[maxIdx] < nums[i]) {
                maxIdx = i;
            }
        }
        return maxIdx;
    }
    private TreeNode constructMaximumBinaryTree(int[] nums, int l, int r) {
        if (l == r) return null;

        int maxIdx = findMaxIndex(nums, l, r);
        TreeNode root = new TreeNode(nums[maxIdx]);
        root.left = constructMaximumBinaryTree(nums, l, maxIdx);
        root.right = constructMaximumBinaryTree(nums, maxIdx + 1, r);
        return root;
    }
    public TreeNode constructMaximumBinaryTree(int[] nums) {
        return constructMaximumBinaryTree(nums, 0, nums.length);
    }
}
// @lc code=end

