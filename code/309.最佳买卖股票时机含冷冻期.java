/*
 * @lc app=leetcode.cn id=309 lang=java
 *
 * [309] 最佳买卖股票时机含冷冻期
 *
 * https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (57.09%)
 * Likes:    583
 * Dislikes: 0
 * Total Accepted:    61.2K
 * Total Submissions: 107K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * 给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
 * 
 * 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
 * 
 * 
 * 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
 * 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
 * 
 * 
 * 示例:
 * 
 * 输入: [1,2,3,0,2]
 * 输出: 3 
 * 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
 * 
 */

// @lc code=start
// dp
// class Solution {
//     public int maxProfit(int[] prices) {
//         if (prices == null || prices.length < 2) return 0;

//         int hold[] = new int[prices.length];
//         int unhold[] = new int[prices.length];

//         // base case
//         hold[0] = -prices[0];
//         hold[1] = Math.max(hold[0], -prices[1]);
//         unhold[0] = 0;
//         unhold[1] = Math.max(unhold[0], hold[0] + prices[1]);

//         for (int i = 2; i < prices.length; i++) {
//             hold[i] = Math.max(hold[i - 1], unhold[i - 2] - prices[i]);
//             unhold[i] = Math.max(unhold[i - 1], hold[i - 1] + prices[i]);
//         }
//         return unhold[prices.length - 1];

//     }
// }

// state machine
class Solution {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length < 2) return 0;

        //state
        // no sold by today
        int rest = 0;
        // sold by today
        int sold = 0;
        // hold
        int hold = Integer.MIN_VALUE;

        for (int price : prices) {
            int preRest = rest;

            rest = Math.max(sold, rest);
            sold = hold + price;
            hold = Math.max(hold, preRest - price);
        }
        return Math.max(rest, sold);
    }
}
// @lc code=end

