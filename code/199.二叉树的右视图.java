import java.util.Deque;
import java.util.LinkedList;
import java.util.List;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=199 lang=java
 *
 * [199] 二叉树的右视图
 *
 * https://leetcode-cn.com/problems/binary-tree-right-side-view/description/
 *
 * algorithms
 * Medium (64.56%)
 * Likes:    370
 * Dislikes: 0
 * Total Accepted:    76.2K
 * Total Submissions: 118.1K
 * Testcase Example:  '[1,2,3,null,5,null,4]'
 *
 * 给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
 * 
 * 示例:
 * 
 * 输入: [1,2,3,null,5,null,4]
 * 输出: [1, 3, 4]
 * 解释:
 * 
 * ⁠  1            <---
 * ⁠/   \
 * 2     3         <---
 * ⁠\     \
 * ⁠ 5     4       <---
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
//     private TreeNode addTree(Deque<TreeNode> stack) {
//         TreeNode tree = stack.poll();
//         if (tree.left != null) {
//             stack.add(tree.left);
//         }
//         if (tree.right != null) {
//             stack.add(tree.right);
//         }
//         return tree;
//     }
//     public List<Integer> rightSideView(TreeNode root) {
//         List<Integer> ans = new LinkedList<>();
//         Deque<TreeNode> stack = new LinkedList<>();
//         if (root == null) return ans;
//         stack.add(root);
//         while (!stack.isEmpty()) {
//             int size = stack.size();
//             int count = 0;
//             while (count++ < size - 1 ) {
//                 addTree(stack);
//             }
//             TreeNode temp = addTree(stack);
//             ans.add(temp.val);
//         }
//         return ans;
//     }
// }
class Solution {
    List<Integer> ans;
    private void dfs(TreeNode root,int depth) {
        if (root == null) return;

        if (depth == ans.size()) {
            ans.add(root.val);
        }
        depth++;
        dfs(root.right, depth);
        dfs(root.left, depth);
    }
    public List<Integer> rightSideView(TreeNode root) {
        ans = new LinkedList<Integer>();
        dfs(root, 0);
        return ans;
    }
}
// @lc code=end

