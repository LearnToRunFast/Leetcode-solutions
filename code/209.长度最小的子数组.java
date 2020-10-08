/*
 * @lc app=leetcode.cn id=209 lang=java
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (44.38%)
 * Likes:    466
 * Dislikes: 0
 * Total Accepted:    91.2K
 * Total Submissions: 205.3K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续
 * 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 * 
 * 
 */

// @lc code=start
class Solution {
    public int minSubArrayLen(int s, int[] nums) {

        int l = 0, h = 0, min = Integer.MAX_VALUE;
        int sum = 0;
        while (h < nums.length) {
            sum += nums[h++];
            while (sum >= s) {
                min = Math.min(min, h - l);
                sum -= nums[l++];
            }
        }
        return min == Integer.MAX_VALUE ? 0 : min;
    }
}
// @lc code=end

