import java.util.ArrayList;
import java.util.List;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=230 lang=java
 *
 * [230] 二叉搜索树中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/description/
 *
 * algorithms
 * Medium (71.62%)
 * Likes:    297
 * Dislikes: 0
 * Total Accepted:    72.8K
 * Total Submissions: 101K
 * Testcase Example:  '[3,1,4,null,2]\n1'
 *
 * 给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
 * 
 * 说明：
 * 你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
 * 
 * 示例 1:
 * 
 * 输入: root = [3,1,4,null,2], k = 1
 * ⁠  3
 * ⁠ / \
 * ⁠1   4
 * ⁠ \
 * 2
 * 输出: 1
 * 
 * 示例 2:
 * 
 * 输入: root = [5,3,6,2,4,null,null,1], k = 3
 * ⁠      5
 * ⁠     / \
 * ⁠    3   6
 * ⁠   / \
 * ⁠  2   4
 * ⁠ /
 * ⁠1
 * 输出: 3
 * 
 * 进阶：
 * 如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
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

 // one query version
class Solution {
    private int count = 0;
    private int ans = 0;

    private void dfs(TreeNode root, int k) {

        if (root == null) return;

        dfs(root.left, k);
        count++;
        if (count == k) {
            ans = root.val;
            return;
        }

        dfs(root.right, k);
        return;
    }

    public int kthSmallest(TreeNode root, int k) {
        dfs(root, k);
        return ans;
    }
}


//  // multiple queries version
//  class Solution {

//     private List<Integer> dfs(TreeNode root, int k, List<Integer> nums, int i) {

//         if (root == null) return nums;

//         dfs(root.left, k, nums, i);
//         i++;
//         nums.add(root.val);
//         if (i == k) return nums;
        
//         dfs(root.right, k, nums, i);
//         return nums;
//     }

//     public int kthSmallest(TreeNode root, int k) {
//         List<Integer> nums = new ArrayList<>();
//         dfs(root, k, nums, 0);
//         return nums.get(k - 1);
//     }
// }
// @lc code=end

