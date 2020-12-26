import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=653 lang=java
 *
 * [653] 两数之和 IV - 输入 BST
 *
 * https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (57.11%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    25.3K
 * Total Submissions: 44.2K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 * 
 * 案例 1:
 * 
 * 
 * 输入: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 9
 * 
 * 输出: True
 * 
 * 
 * 
 * 
 * 案例 2:
 * 
 * 
 * 输入: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 28
 * 
 * 输出: False
 * 
 * 
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
// class Solution {
    
//     private boolean dfs(TreeNode root, int k, Set<Integer> numSet) {
//         if (root == null) return false;

//         if (numSet.contains(k - root.val)) return true;
//         numSet.add(root.val);
//         return dfs(root.left, k, numSet) || dfs(root.right, k, numSet);
        
//     }
//     public boolean findTarget(TreeNode root, int k) {
//         Set<Integer> numSet = new HashSet<>();
//         return dfs(root, k, numSet);
//     }
// }
class Solution {
    
    private void inorder(TreeNode root, List<Integer> list) {
        if (root == null) return;
        inorder(root.left, list);
        list.add(root.val);
        inorder(root.right, list);
    }
    public boolean findTarget(TreeNode root, int k) {
        List<Integer> list = new ArrayList<>();
        inorder(root, list);

        int lo = 0, hi = list.size() - 1;
        while (lo < hi) {
            int sum = list.get(lo) + list.get(hi);
            if (sum == k) return true;
            if (sum > k) hi--;
            else lo++;
        }
        return false;
    }
}
// @lc code=end

