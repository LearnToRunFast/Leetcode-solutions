import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=322 lang=java
 *
 * [322] 零钱兑换
 *
 * https://leetcode-cn.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (41.03%)
 * Likes:    883
 * Dislikes: 0
 * Total Accepted:    147.1K
 * Total Submissions: 352.6K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * 给定不同面额的硬币 coins 和一个总金额
 * amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
 * 
 * 你可以认为每种硬币的数量是无限的。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：coins = [1, 2, 5], amount = 11
 * 输出：3 
 * 解释：11 = 5 + 5 + 1
 * 
 * 示例 2：
 * 
 * 
 * 输入：coins = [2], amount = 3
 * 输出：-1
 * 
 * 示例 3：
 * 
 * 
 * 输入：coins = [1], amount = 0
 * 输出：0
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：coins = [1], amount = 1
 * 输出：1
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：coins = [1], amount = 2
 * 输出：2
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
// dp memo
// class Solution {
//     private int coinChange(int[] coins, int amount, int[] memo) {
//         if (amount == 0) return 0;
//         if (amount < 0 ) return -1;

//         if (memo[amount] != 0) {
//             return memo[amount];
//         }
        
//         int min = Integer.MAX_VALUE;
//         for (int coin : coins) {
//             if (coin > amount) break;
//             int ans = coinChange(coins, amount - coin, memo);
//             if (ans >= 0) {
//                 ans++;
//                 min = Math.min(min, ans);
//             }
//         }
//         memo[amount] = min == Integer.MAX_VALUE ? -1 : min;
//         return memo[amount];

//     }
//     public int coinChange(int[] coins, int amount) {
//         int[] memo = new int[amount + 1];
//         Arrays.sort(coins);
//         return coinChange(coins, amount, memo);
//     }
// }

class Solution {
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        int max = amount + 1;
        Arrays.fill(dp, max);
        dp[0] = 0;
        for (int i = 1; i <= amount; i++) {
            for(int coin: coins) {
                if (i >= coin) {
                    dp[i] = Math.min(dp[i], dp[i - coin] + 1);
                }
            }
        }
        return dp[amount] == max ? -1 : dp[amount];

    }
}
// @lc code=end

