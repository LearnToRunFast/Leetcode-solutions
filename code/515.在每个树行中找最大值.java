import java.awt.List;
import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=515 lang=java
 *
 * [515] 在每个树行中找最大值
 *
 * https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/description/
 *
 * algorithms
 * Medium (62.74%)
 * Likes:    115
 * Dislikes: 0
 * Total Accepted:    22.1K
 * Total Submissions: 35.1K
 * Testcase Example:  '[1,3,2,5,3,null,9]'
 *
 * 您需要在二叉树的每一行中找到最大的值。
 * 
 * 示例：
 * 
 * 
 * 输入: 
 * 
 * ⁠         1
 * ⁠        / \
 * ⁠       3   2
 * ⁠      / \   \  
 * ⁠     5   3   9 
 * 
 * 输出: [1, 3, 9]
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
//dfs
class Solution {
    private void dfs(TreeNode root, List<Integer> ans, int level) {
        if (root == null) return;
        if (ans.size() <= level) {
            ans.add(root.val);
        } else {
            ans.set(level, Math.max(ans.get(level), root.val));
        }
        dfs(root.left, ans, level + 1);
        dfs(root.right, ans, level + 1);

    }
    public List<Integer> largestValues(TreeNode root) {
        List<Integer> ans = new ArrayList<>();
        dfs(root, ans, 0);
        return ans;
    }
}
// @lc code=end

