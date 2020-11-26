import java.util.ArrayList;
import java.util.LinkedList;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=102 lang=java
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.69%)
 * Likes:    703
 * Dislikes: 0
 * Total Accepted:    220.8K
 * Total Submissions: 346.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 
 * 
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层次遍历结果：
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode(int x) { val = x; } }
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        LinkedList<TreeNode> stack = new LinkedList<>();
        List<List<Integer>> ans = new ArrayList<>();
        if (root == null)
            return ans;
        stack.add(root);
        while (!stack.isEmpty()) {
            int level = stack.size();
            List<Integer> subList = new ArrayList();
            for (int i = 0; i < level; i++) {
                TreeNode curr = stack.pop();
                if (curr.left != null) {
                    stack.add(curr.left);
                }
                if (curr.right != null) {
                    stack.add(curr.right);
                }
                subList.add(curr.val);
            }
            ans.add(subList);
        }
        return ans;
    }
}
// @lc code=end
