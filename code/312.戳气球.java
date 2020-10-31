import java.lang.reflect.Array;
import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=312 lang=java
 *
 * [312] 戳气球
 *
 * https://leetcode-cn.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (66.97%)
 * Likes:    562
 * Dislikes: 0
 * Total Accepted:    33.1K
 * Total Submissions: 49.3K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 * 
 * 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的
 * left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
 * 
 * 求所能获得硬币的最大数量。
 * 
 * 说明:
 * 
 * 
 * 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
 * 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
 * 
 * 
 * 示例:
 * 
 * 输入: [3,1,5,8]
 * 输出: 167 
 * 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
 * coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167
 * 
 * 
 */

// @lc code=start
class Solution {
    public int maxCoins(int[] nums) {
        int n = nums.length;
        int size = n + 2;
        int[] points = new int[size];
        
        points[0] = 1;
        points[n + 1] = 1;
        for (int i = 1; i <= n; i++) {
            points[i] = nums[i - 1];
        }


        // now we use points instead of nums

        //dp[i][j] = dp[i][k] + dp[k][j] + points[i]*points[k]*points[j]
        // where k is the last element between i and j
        int[][] dp = new int[size][size];

        for (int i = n; i >= 0; i--) {
            for (int j = i + 2; j < size; j++) {
                for (int k = i + 1; k < j; k++) {
                    dp[i][j] = Math.max(
                        dp[i][j], 
                        dp[i][k] + points[i] * points[k] * points[j] + dp[k][j]);
                }
            }
        }
        return dp[0][n + 1];
    }
}
// @lc code=end

