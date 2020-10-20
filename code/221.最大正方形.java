import java.math.BigInteger;

/*
 * @lc app=leetcode.cn id=221 lang=java
 *
 * [221] 最大正方形
 *
 * https://leetcode-cn.com/problems/maximal-square/description/
 *
 * algorithms
 * Medium (42.78%)
 * Likes:    590
 * Dislikes: 0
 * Total Accepted:    76.4K
 * Total Submissions: 177.6K
 * Testcase Example:  '[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]'
 *
 * 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
 * 
 * 示例:
 * 
 * 输入: 
 * 
 * 1 0 1 0 0
 * 1 0 1 1 1
 * 1 1 1 1 1
 * 1 0 0 1 0
 * 
 * 输出: 4
 * 
 */

// @lc code=start
class Solution {
    public int maximalSquare(char[][] matrix) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) return 0;

        int m = matrix.length;
        int n = matrix[0].length;
        int maxLen = 0;
        // int[][] dp = new int[m + 1][n + 1];
        // for (int i = 0; i < m; i++) {
        //     for (int j = 0; j < n; j++) {
        //         if (matrix[i][j] == '1') {
        //             dp[i + 1][j + 1] = 1 + Math.min(dp[i][j + 1],
        //                                     Math.min(dp[i + 1][j], dp[i][j]));
        //         }
        //         maxLen = Math.max(maxLen, dp[i + 1][j + 1]);
        //     }
        // }

        int[] dp = new int[n + 1];
        for (char[] col : matrix) {
            int northWest = 0;
            for (int i = 0; i < n; i++) {
                int nextNorthWest = dp[i + 1];
                if (col[i] == '1') {
                    dp[i + 1] = 1 + Math.min(northWest, Math.min(dp[i], dp[i + 1]));
                    maxLen = Math.max(maxLen, dp[i + 1]);
                } else {
                    dp[i + 1] = 0;
                }
                northWest = nextNorthWest;
            }
        }
        return maxLen * maxLen;
    }
}

// @lc code=end
/*
1 1 1 1 0 
1 1 1 1 0 
1 1 1 1 1 
1 1 1 1 1 
0 0 1 1 1 */
