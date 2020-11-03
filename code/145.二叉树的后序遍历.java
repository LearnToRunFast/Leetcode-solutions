import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=145 lang=java
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (72.27%)
 * Likes:    471
 * Dislikes: 0
 * Total Accepted:    159K
 * Total Submissions: 216.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 * 
 * 示例:
 * 
 * 输入: [1,null,2,3]  
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3 
 * 
 * 输出: [3,2,1]
 * 
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
// recursion
// class Solution {
//     private void postorder(TreeNode root, List<Integer> ans) {
//         if (root == null) return;

//         postorder(root.left, ans);
//         postorder(root.right, ans);
//         ans.add(root.val);
//     }
//     public List<Integer> postorderTraversal(TreeNode root) {
//         List<Integer> ans = new ArrayList<>();
//         postorder(root, ans);
//         return ans;
//     }
// }

// use preorder + reverse
// class Solution {

//     public List<Integer> postorderTraversal(TreeNode root) {
//         LinkedList<TreeNode> stack = new LinkedList<>();
//         List<Integer> ans = new ArrayList<>();

//         while (root != null || !stack.isEmpty()) {
//             while (root != null) {
//                 ans.add(root.val);
//                 stack.push(root.left);
//                 root = root.right;
//             }
//             root = stack.pop();
//         }

//         Collections.reverse(ans);
//         return ans;
//     }
// }
// // v2
// class Solution {

//     public List<Integer> postorderTraversal(TreeNode root) {
//         List<Integer> ans = new ArrayList<>();
//         if (root == null) return ans;

//         LinkedList<TreeNode> stack = new LinkedList<>();
//         TreeNode mark = root;
//         stack.push(root);
//         while (!stack.isEmpty()) {
//             TreeNode curr = stack.peek();
//             if (curr.left != null && curr.left != mark && curr.right != mark) {
//                 stack.push(curr.left);
//             } else if (curr.right != null && curr.right != mark) {
//                 stack.push(curr.right);
//             } else {
//                 mark = stack.pop();
//                 ans.add(mark.val);
//             }
//         }

//         return ans;
//     }
// }
// v3
class Solution {

    public List<Integer> postorderTraversal(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        LinkedList<TreeNode> stack = new LinkedList<>();
        
        TreeNode prev = null;
        while (root != null || !stack.isEmpty()) {
            while (root != null) {
                stack.push(root);
                root = root.left;
            }

            root = stack.peek();
            if (root.right == null || root.right == prev) {
                ans.add(stack.pop().val);
                prev = root;
                root = null;
            } else {
                root = root.right;
            }

        }
        return ans;
    }
}
// @lc code=end

