import java.util.ArrayList;

import apple.laf.JRSUIUtils.Tree;

/*
 * @lc app=leetcode.cn id=95 lang=java
 *
 * [95] 不同的二叉搜索树 II
 *
 * https://leetcode-cn.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (66.82%)
 * Likes:    706
 * Dislikes: 0
 * Total Accepted:    66.2K
 * Total Submissions: 99.1K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：3
 * 输出：
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * 解释：
 * 以上的输出对应以下 5 种不同结构的二叉搜索树：
 * 
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= n <= 8
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
// recursion
// class Solution {
//     private List<TreeNode> generateTrees(int start, int end) {
//         List<TreeNode> ans = new ArrayList();

//         if (start > end) {
//             ans.add(null);
//             return ans;    
//         }

//         if (start == end) {
//             TreeNode tree = new TreeNode(start);
//             ans.add(tree);
//             return ans;
//         }

//         for (int i = start; i <= end; i++) {
//             List<TreeNode> leftTrees = generateTrees(start, i - 1);
//             List<TreeNode> rightTrees = generateTrees(i + 1, end);
            
//             for (TreeNode leftTree : leftTrees) {
//                 for (TreeNode rightTree : rightTrees) {
//                     TreeNode tree = new TreeNode(i);
//                     tree.left = leftTree;
//                     tree.right = rightTree;
//                     ans.add(tree);
//                 }
//             }
//         }
//         return ans;
//     }
//     public List<TreeNode> generateTrees(int n) {
//         if (n < 1) return new ArrayList();

//         return generateTrees(1, n);
//     }
// }
// DP
// class Solution {

//     private TreeNode clone(TreeNode n, int offset) {
//         if (n == null) return null;
//         TreeNode tree = new TreeNode(n.val + offset);
//         tree.left = clone(n.left, offset);
//         tree.right = clone(n.right, offset);
//         return tree;
//     }

//     public List<TreeNode> generateTrees(int n) {
//         List<TreeNode>[] dp = new ArrayList[n + 1];
//         dp[0] = new ArrayList<TreeNode>();
//         if (n == 0) {
//             return dp[0];
//         }
//         dp[0].add(null);

//         for (int len = 1; len <= n; len++) {
//             dp[len] = new ArrayList<>();
//             for (int root = 1; root <= len; root++) {
//                 int left = root - 1;
//                 int right = len - root;
//                 for (TreeNode leftTree : dp[left]) {
//                     for (TreeNode rightTree : dp[right]) {
//                         TreeNode treeRoot = new TreeNode(root);
//                         treeRoot.left = leftTree;
//                         treeRoot.right = clone(rightTree, root);
//                         dp[len].add(treeRoot);
//                     }
//                 }
//             }
//         }
//         return dp[n];
//     }
// }
// dp 2
class Solution {

    private TreeNode clone(TreeNode root) {
        if (root == null) return null;
        TreeNode tree = new TreeNode(root.val);
        tree.left = clone(root.left);
        tree.right = clone(root.right);
        return tree;
    }

    public List<TreeNode> generateTrees(int n) {
        List<TreeNode> pre = new ArrayList<>();
        if (n == 0) {
            return pre;
        }
        pre.add(null);

        for (int i = 1; i <= n; i++) {
            List<TreeNode> ans = new ArrayList<>();
            for (TreeNode root : pre) {
                // ans as root
                TreeNode tree = new TreeNode(i);
                tree.left = root;
                ans.add(tree);
                
                // as right child
                for (int j = 0; j <= i; j++) {
                    TreeNode rootCopy = clone(root);
                    TreeNode curr = rootCopy;
                    
                    int step = 0;
                    while (curr != null && step++ < j) {
                        curr = curr.right;
                    }

                    if (curr == null) break;

                    tree = new TreeNode(i);
                    TreeNode tmpTree = curr.right;
                    tree.left = tmpTree;
                    curr.right = tree;
                    ans.add(rootCopy);
                }
            }
            pre = ans;
        }
        return pre;
    }
}
// @lc code=end

