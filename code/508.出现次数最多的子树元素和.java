import java.awt.List;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

/*
 * @lc app=leetcode.cn id=508 lang=java
 *
 * [508] 出现次数最多的子树元素和
 *
 * https://leetcode-cn.com/problems/most-frequent-subtree-sum/description/
 *
 * algorithms
 * Medium (65.52%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    9.8K
 * Total Submissions: 14.9K
 * Testcase Example:  '[5,2,-3]'
 *
 * 给你一个二叉树的根结点，请你找出出现次数最多的子树元素和。一个结点的「子树元素和」定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
 * 
 * 你需要返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
 * 
 * 
 * 
 * 示例 1：
 * 输入:
 * 
 * ⁠ 5
 * ⁠/  \
 * 2   -3
 * 
 * 
 * 返回 [2, -3, 4]，所有的值均只出现一次，以任意顺序返回所有值。
 * 
 * 示例 2：
 * 输入：
 * 
 * ⁠ 5
 * ⁠/  \
 * 2   -5
 * 
 * 
 * 返回 [2]，只有 2 出现两次，-5 只出现 1 次。
 * 
 * 
 * 
 * 提示： 假设任意子树元素和均可以用 32 位有符号整数表示。
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
    private int max = 0;
    private int dfs(TreeNode root, Map<Integer, Integer> map) {
        if (root == null) {
            return 0;
        }
        int left = dfs(root.left, map);
        int right = dfs(root.right, map);
        int sum = root.val + left + right;
        map.put(sum, map.getOrDefault(sum, 0) + 1);
        max = Math.max(max, map.get(sum));

        return sum;
    }
    public int[] findFrequentTreeSum(TreeNode root) {
        Map<Integer, Integer> map = new HashMap<>();
        dfs(root, map);
        List<Integer> list = new ArrayList<>();
        for (int key : map.keySet()) {
            if (max == map.get(key)) {
                list.add(key);
            }
        }
        int[] ans = new int[list.size()];
        for (int i = 0; i < list.size(); i++) {
            ans[i] = list.get(i);
        }
        return ans;
    }
}
// @lc code=end

