/*
 * @lc app=leetcode.cn id=132 lang=java
 *
 * [132] 分割回文串 II
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning-ii/description/
 *
 * algorithms
 * Hard (44.61%)
 * Likes:    237
 * Dislikes: 0
 * Total Accepted:    19.4K
 * Total Submissions: 43.4K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 * 
 * 返回符合要求的最少分割次数。
 * 
 * 示例:
 * 
 * 输入: "aab"
 * 输出: 1
 * 解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
 * 
 * 
 */

// @lc code=start
class Solution {
    // define dp[i] the minCut of [0, i]
    // we cut [0, i] in point j
    // if [j + 1, i] is palindrone
    // then dp[i] = dp[j] + 1;
    //  we need to find min for j in [0, i]
    // dp[i] = min(dp[i], dp[j] + 1)
    public int minCut(String s) {
        int[] dp = new int[s.length()];
        // max the value in dp[i]
        for (int i = 1; i < dp.length; i++) {
            dp[i] = i;
        }

        // set up 2d dp for palindrone check
        boolean[][] isPalindrone = new boolean[s.length()][s.length()];
        for (int r = 0; r < s.length(); r++) {
            for (int l = 0; l <= r; l++) {
                if (s.charAt(l) == s.charAt(r) &&
                    ((r - l <= 2) || isPalindrone[l + 1][r - 1])) {
                        isPalindrone[l][r] = true;
                    }
            }
        }

        for (int i = 1; i < s.length(); i++ ) {
            if (isPalindrone[0][i]) {
                dp[i] = 0;
                continue;
            }

            for (int j = 0; j < i; j++) {
                if (isPalindrone[j + 1][i]) {
                    dp[i] = Math.min(dp[i], dp[j] + 1);
                }
            }
        }
        return dp[s.length() - 1];

    }
}
// @lc code=end

