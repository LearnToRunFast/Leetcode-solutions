/*
 * @lc app=leetcode.cn id=5 lang=java
 *
 * [5] 最长回文子串
 *
 * https://leetcode-cn.com/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (32.46%)
 * Likes:    3072
 * Dislikes: 0
 * Total Accepted:    451.3K
 * Total Submissions: 1.4M
 * Testcase Example:  '"babad"'
 *
 * 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
 * 
 * 示例 1：
 * 
 * 输入: "babad"
 * 输出: "bab"
 * 注意: "aba" 也是一个有效答案。
 * 
 * 
 * 示例 2：
 * 
 * 输入: "cbbd"
 * 输出: "bb"
 * 
 * 
 */

// @lc code=start
class Solution {
    // private boolean isPalindrome(int l, int r, int[] dp) {
    //     // convert to newS
    //     l = l * 2 + 1;
    //     r = r * 2 + 1;
    //     int mid = l + (r - l) / 2;
    //     if (mid >= 0 && mid < dp.length && r - l + 1 <= dp[mid]) {
    //         return true;
    //     }
    //     return false;
    // }
    public String longestPalindrome(String s) {
        // manacher
        char[] newS = new char[s.length() * 2 + 1];
        for (int i = 0; i < newS.length; i++) {
            newS[i] = (i & 1) == 1 ? s.charAt(i >> 1) : '0'; 
        }

        int l = 0, r = -1;
        int[] dp = new int[newS.length];

        for (int i = 0; i < newS.length; i++) {
            // in range
            if (i <= r) {
                dp[i] = Math.min((r - i) * 2 + 1, dp[l + r - i]);
            } else {
                dp[i] = 1;
            }

            int temp_l = i - dp[i] / 2 - 1;
            int temp_r = i + dp[i] / 2 + 1;

            while (temp_l >= 0 && temp_r < newS.length && newS[temp_l] == newS[temp_r]) {
                temp_l--;
                temp_r++;
            }
            dp[i] = temp_r - temp_l - 1;
            if (temp_r > r) {
                l = temp_l + 1;
                r = temp_r - 1;
            }
        }

        
        // find longest and pos
        int pos = -1, len = -1;
        for (int i = 0; i < dp.length; i++) {
            if(dp[i] > len) {
                len = dp[i];
                pos = i;
            }
        }
        // for(int i = 0; i < dp.length; i++) {
        //     System.out.printf("%s ",newS[i]);
        // }
        // System.out.println();
        // for(int i = 0; i < dp.length; i++) {
        //     System.out.printf("%d ",dp[i]);
        // }
        // System.out.println();

        // System.out.println("len " + len + " pos " + pos);
        
        // map back
        pos = pos / 2 - len / 4;
        len = len / 2;
        return s.substring(pos, pos + len);


    }
}
// @lc code=end

