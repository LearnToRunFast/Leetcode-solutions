/*
 * @lc app=leetcode.cn id=264 lang=java
 *
 * [264] 丑数 II
 *
 * https://leetcode-cn.com/problems/ugly-number-ii/description/
 *
 * algorithms
 * Medium (54.98%)
 * Likes:    461
 * Dislikes: 0
 * Total Accepted:    43.2K
 * Total Submissions: 78.3K
 * Testcase Example:  '10'
 *
 * 编写一个程序，找出第 n 个丑数。
 * 
 * 丑数就是质因数只包含 2, 3, 5 的正整数。
 * 
 * 示例:
 * 
 * 输入: n = 10
 * 输出: 12
 * 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
 * 
 * 说明:  
 * 
 * 
 * 1 是丑数。
 * n 不超过1690。
 * 
 * 
 */

// @lc code=start
class Helper {
    private static final int R = 1690;
    private static int[] dp;
    private static int i;
    private static int p2 = 0, p3 = 0, p5 = 0;
    Helper() {
        i = 1;
        dp = new int[R];
        dp[0] = 1;
    }
    public int getUgly(int n) {
        for (;i < n;i++) {
            int n2 = dp[p2] * 2, n3 = dp[p3] * 3, n5 = dp[p5] * 5;
            dp[i] = Math.min(n2, Math.min(n3, n5));
            if (dp[i] == n2) p2++;
            if (dp[i] == n3) p3++;
            if (dp[i] == n5) p5++;
        }
        return dp[n - 1];
    }
}
class Solution {
    public static Helper ugly = new Helper();
    public int nthUglyNumber(int n) {
        return ugly.getUgly(n);
    }
}
// @lc code=end

