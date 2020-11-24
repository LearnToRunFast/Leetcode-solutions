import java.util.LinkedList;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=98 lang=java
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.75%)
 * Likes:    834
 * Dislikes: 0
 * Total Accepted:    192.8K
 * Total Submissions: 587.9K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 * 
 * 假设一个二叉搜索树具有如下特征：
 * 
 * 
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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
// in order traversal
// class Solution {

//     public boolean isValidBST(TreeNode root) {
//         LinkedList<TreeNode> stack = new LinkedList<>();
//         Integer prev = null;
//         while (root != null || !stack.isEmpty()) {
//             while (root != null) {
//                 stack.push(root);
//                 root = root.left;
//             }
//             TreeNode curr = stack.pop();
//             if (prev != null && prev >= curr.val) {
//                 return false;
//             }
//             prev = curr.val;
//             root = curr.right;
//         }
//         return true;
//     }
// }
class Solution {
    private boolean isValidBST(TreeNode root, Integer lower, Integer upper) {
        if (root == null) return true;

        int val = root.val;

        if (lower != null && val <= lower) return false;
        if (upper != null && val >= upper) return false;

        return isValidBST(root.left, lower, val) && isValidBST(root.right, val, upper);
    }
    public boolean isValidBST(TreeNode root) {
        return isValidBST(root, null, null);
    }
}
// @lc code=end

