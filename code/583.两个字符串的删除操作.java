/*
 * @lc app=leetcode.cn id=583 lang=java
 *
 * [583] 两个字符串的删除操作
 *
 * https://leetcode-cn.com/problems/delete-operation-for-two-strings/description/
 *
 * algorithms
 * Medium (51.78%)
 * Likes:    172
 * Dislikes: 0
 * Total Accepted:    13.4K
 * Total Submissions: 25.8K
 * Testcase Example:  '"sea"\n"eat"'
 *
 * 给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入: "sea", "eat"
 * 输出: 2
 * 解释: 第一步将"sea"变为"ea"，第二步将"eat"变为"ea"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给定单词的长度不超过500。
 * 给定单词中的字符只含有小写字母。
 * 
 * 
 */

// @lc code=start
// class Solution {
//     public int minDistance(String word1, String word2) {
//         int m = word1.length() + 1;
//         int n = word2.length() + 1;
//         int[][] dp = new int[m][n];
//         for (int i = 1; i < m; i++) {
//             for (int j = 1; j < n; j++) {
//                 if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
//                     dp[i][j] = dp[i - 1] [j - 1] + 1;
//                 } else {
//                     dp[i][j] = Math.max(dp[i][j - 1], dp[i - 1][j]);
//                 }
//             }
//         }
//         // now dp[m - 1][ n - 1] is longest common sequences.
//         return word1.length() + word2.length() - 2 * dp[m - 1][n - 1];
//     }
// }
// normal 2-d dp
// class Solution {
//     public int minDistance(String word1, String word2) {
//         int m = word1.length() + 1;
//         int n = word2.length() + 1;
//         int[][] dp = new int[m][n];
//         for (int i = 0; i < m; i++) {
//             dp[i][0] = i;
//         }
//         for (int i = 0; i < n; i++) {
//             dp[0][i] = i;
//         }
//         for (int i = 1; i < m; i++) {
//             for (int j = 1; j < n; j++) {
//                 if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
//                     dp[i][j] = dp[i - 1] [j - 1];
//                 } else {
//                     dp[i][j] = 1 + Math.min(dp[i][j - 1], dp[i - 1][j]);
//                 }
//             }
//         }
//         return dp[m - 1][n - 1];
//     }
// }
class Solution {
    public int minDistance(String word1, String word2) {
        int[] dp = new int[word2.length() + 1];
        for (int i = 0; i <= word1.length(); i++) {
            int[] temp = new int[word2.length() + 1];
            for (int j = 0; j <= word2.length(); j++) {
                if (i == 0 || j == 0) {
                    temp[j] = i + j;
                } else if (word1.charAt(i - 1)  == word2.charAt(j - 1)) {
                    temp[j] = dp[j - 1];
                } else {
                    temp[j] = 1 + Math.min(dp[j], temp[j - 1]);
                }
            }
            dp = temp;
        }
        return dp[word2.length()];
    }
}
// @lc code=end

