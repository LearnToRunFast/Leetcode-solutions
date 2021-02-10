/*
 * @lc app=leetcode.cn id=416 lang=java
 *
 * [416] 分割等和子集
 *
 * https://leetcode-cn.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (49.18%)
 * Likes:    671
 * Dislikes: 0
 * Total Accepted:    99.5K
 * Total Submissions: 201.7K
 * Testcase Example:  '[1,5,11,5]'
 *
 * 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
 * 
 * 注意:
 * 
 * 
 * 每个数组中的元素不会超过 100
 * 数组的大小不会超过 200
 * 
 * 
 * 示例 1:
 * 
 * 输入: [1, 5, 11, 5]
 * 
 * 输出: true
 * 
 * 解释: 数组可以分割成 [1, 5, 5] 和 [11].
 * 
 * 
 * 
 * 
 * 示例 2:
 * 
 * 输入: [1, 2, 3, 5]
 * 
 * 输出: false
 * 
 * 解释: 数组不能分割成两个元素和相等的子集.
 * 
 * 
 * 
 * 
 */

// @lc code=start
class Solution {
    public boolean canPartition(int[] nums) {
        // dp[w] define max value of weight w
        int sum = 0;
        for (int num : nums) {
            sum += num;
        }
        if ((sum & 1) == 1) return false;
        int target = sum >> 1;
        int[] dp = new int[target + 1];
        for (int i = 0; i < nums.length; i++) {
            for (int w = target; w >= nums[i]; w--) {
                dp[w] = Math.max(dp[w], dp[w - nums[i]] + nums[i]);
            }
        }
        return dp[target] == target;
    }
}
// @lc code=end

