/*
 * @lc app=leetcode.cn id=279 lang=java
 *
 * [279] 完全平方数
 *
 * https://leetcode-cn.com/problems/perfect-squares/description/
 *
 * algorithms
 * Medium (57.53%)
 * Likes:    640
 * Dislikes: 0
 * Total Accepted:    92.7K
 * Total Submissions: 159.4K
 * Testcase Example:  '12'
 *
 * 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
 * 
 * 示例 1:
 * 
 * 输入: n = 12
 * 输出: 3 
 * 解释: 12 = 4 + 4 + 4.
 * 
 * 示例 2:
 * 
 * 输入: n = 13
 * 输出: 2
 * 解释: 13 = 4 + 9.
 * 
 */

// @lc code=start
// memo
// class Solution {
//     private int numSquares(int[] memo, int n) {
//         if (memo[n] > 0) return memo[n];

//         int sqRoot = (int) Math.sqrt(n);
//         if (sqRoot << 1 == n) {
//             memo[n] = 1;
//         }
//         int res = n;
//         for (int i = 1; i <= sqRoot; i++) {
//             res = Math.min(res, numSquares(memo, n - i * i) + 1);
//         }
//         memo[n] = res;
//         return memo[n];
//     }
    
//     public int numSquares(int n) {
//         int[] memo = new int[n + 1];
//         return numSquares(memo, n);
//     }
// }
// dp bootm-up
class Solution {
    public int numSquares(int n) {
        int[] dp = new int[n + 1];

        for (int i = 1; i <= n; i++) {
            dp[i]= i;
        }
        
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j * j <= i; j++) {
                dp[i] = Math.min(dp[i], dp[i - j * j] + 1);
            }
        }

        return dp[n];
    }
}
// @lc code=end

