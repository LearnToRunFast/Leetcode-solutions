/*
 * @lc app=leetcode.cn id=72 lang=java
 *
 * [72] 编辑距离
 *
 * https://leetcode-cn.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (60.03%)
 * Likes:    1247
 * Dislikes: 0
 * Total Accepted:    92.3K
 * Total Submissions: 153.8K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
 * 
 * 你可以对一个单词进行如下三种操作：
 * 
 * 
 * 插入一个字符
 * 删除一个字符
 * 替换一个字符
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：word1 = "horse", word2 = "ros"
 * 输出：3
 * 解释：
 * horse -> rorse (将 'h' 替换为 'r')
 * rorse -> rose (删除 'r')
 * rose -> ros (删除 'e')
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：word1 = "intention", word2 = "execution"
 * 输出：5
 * 解释：
 * intention -> inention (删除 't')
 * inention -> enention (将 'i' 替换为 'e')
 * enention -> exention (将 'n' 替换为 'x')
 * exention -> exection (将 'n' 替换为 'c')
 * exection -> execution (插入 'u')
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * word1 和 word2 由小写英文字母组成
 * 
 * 
 */

// @lc code=start
class Solution {
    public int minDistance(String word1, String word2) {
        int m = word1.length(), n = word2.length();
        if (m * n == 0) {
            return m + n;
        }
        // only three opeartions
        // insert to A
        // insert to B
        // modify A
        // dp[i][j] means word1 from 0 - i and word2 from 0 - j's minDistance
        int[][] dp = new int[m + 1][n + 1];
        
        for (int i = 0; i <= m; i++) {
            dp[i][0] = i;
        }

        for (int i = 0; i <= n; i++) {
            dp[0][i] = i;
        }

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                int down = dp[i - 1][j]; // insert into word2
                int left = dp[i][j - 1]; // insert into word1
                int down_left = dp[i - 1][j - 1]; // edit word1
                
                // if prev is same, we need to minus 1
                // as if i - 1 and j - 1 are same
                // dp[i - 1][j - 1] == dp[i][j]
                if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
                    down_left--;
                }
                dp[i][j] = 1 + Math.min(down, Math.min(left, down_left));
            }
        }
        return dp[m][n];
    }
}
// @lc code=end

