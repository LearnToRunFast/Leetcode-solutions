import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=257 lang=java
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (66.31%)
 * Likes:    411
 * Dislikes: 0
 * Total Accepted:    87.4K
 * Total Submissions: 131.7K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
class Solution {
    private void dfs(TreeNode root, List<String> ans, String path) {
        if (root == null) return;

        StringBuilder temp = new StringBuilder(path);
        temp.append(Integer.toString(root.val));
        if (root.left == null && root.right == null) {
            ans.add(temp.toString());
            return;
        }

        temp.append("->");
        dfs(root.left, ans, temp.toString());
        dfs(root.right, ans, temp.toString());

    }
    public List<String> binaryTreePaths(TreeNode root) {
        List<String> ans = new LinkedList<>();
        dfs(root, ans, "");
        return ans;
    }
}
// @lc code=end

