import java.awt.List;
import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=113 lang=java
 *
 * [113] 路径总和 II
 *
 * https://leetcode-cn.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (61.15%)
 * Likes:    388
 * Dislikes: 0
 * Total Accepted:    104K
 * Total Submissions: 169.9K
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * 给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 给定如下二叉树，以及目标和 sum = 22，
 * 
 * ⁠             5
 * ⁠            / \
 * ⁠           4   8
 * ⁠          /   / \
 * ⁠         11  13  4
 * ⁠        /  \    / \
 * ⁠       7    2  5   1
 * 
 * 
 * 返回:
 * 
 * [
 * ⁠  [5,4,11,2],
 * ⁠  [5,8,4,5]
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
    private void dfs(TreeNode root, int sum, List<List<Integer>> ans, List<Integer> subAns) {
        if (root == null) {
            return;
        }
        subAns.add(root.val);
        sum -= root.val;
        if (root.left == null && root.right == null && sum == 0) {
            ans.add(new ArrayList<Integer>(subAns));
        }

        dfs(root.left, sum, ans, subAns);
        dfs(root.right, sum, ans, subAns);
        subAns.remove(subAns.size() - 1);
    }

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        List<List<Integer>> ans = new ArrayList<>();
        if (root == null)
            return ans;
        dfs(root, sum, ans, new ArrayList<Integer>());
        return ans;
    }
}
// @lc code=end
