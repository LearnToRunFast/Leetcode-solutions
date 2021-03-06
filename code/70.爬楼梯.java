/*
 * @lc app=leetcode.cn id=70 lang=java
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (51.21%)
 * Likes:    1467
 * Dislikes: 0
 * Total Accepted:    364.7K
 * Total Submissions: 709.3K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 注意：给定 n 是一个正整数。
 * 
 * 示例 1：
 * 
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 * 
 * 示例 2：
 * 
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 * 
 * 
 */

// @lc code=start
// class Solution {
//     public int climbStairs(int n) {
//         int[] dp = new int[n + 1];
//         dp[0] = 1;
//         dp[1] = 1;
//         for (int i = 2; i <= n; i++) {
//             for (int j = 1; j <= 2; j++) {
//                 dp[i] += dp[i - j];
//             }
//         }
//         return dp[n];
//     }
// }
class Solution {
    public int climbStairs(int n) {
        int pre2 = 1, pre1 = 1;
        for (int i = 2; i <= n; i++) {
            int curr = pre2 + pre1;
            pre2 = pre1;
            pre1 = curr;
        }
        return pre1;
    }
}
// @lc code=end

