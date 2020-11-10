/*
 * @lc app=leetcode.cn id=32 lang=java
 *
 * [32] 最长有效括号
 *
 * https://leetcode-cn.com/problems/longest-valid-parentheses/description/
 *
 * algorithms
 * Hard (34.04%)
 * Likes:    1060
 * Dislikes: 0
 * Total Accepted:    111K
 * Total Submissions: 326K
 * Testcase Example:  '"(()"'
 *
 * 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
 * 
 * 示例 1:
 * 
 * 输入: "(()"
 * 输出: 2
 * 解释: 最长有效括号子串为 "()"
 * 
 * 
 * 示例 2:
 * 
 * 输入: ")()())"
 * 输出: 4
 * 解释: 最长有效括号子串为 "()()"
 * 
 * 
 */

// @lc code=start
// dp
class Solution {
    public int longestValidParentheses(String s) {
        int[] dp = new int[s.length()];
        int max = 0;
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i - 1) == '(') {
                    dp[i] = i > 2 ? dp[i - 2] + 2 : 2;
                } else if (i - dp[i - 1] - 1 >= 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
                    dp[i] += dp[i - 1] + 2;
                    dp[i] += i - dp[i - 1] - 2 >= 0 ? dp[ i - dp[i - 1] - 2] : 0;
                }
                max = Math.max(max, dp[i]);
            }
        }
        return max;
    }
}

// Double Pointers
// class Solution {
//     public int longestValidParentheses(String s) {
//         int max = 0;
//         int l = 0, r = 0;
//         for (char c : s.toCharArray()) {
//             if (c == '(') {
//                 l++;
//             } else {
//                 r++;
//             }
//             if (l == r) {
//                 max = Math.max(max, l * 2);
//             } else if (r > l) {
//                 l = r = 0;
//             }
//         }
//         l = r = 0;
//         for (int i = s.length() - 1; i >= 0; i--) {
//             if (s.charAt(i) == '(') {
//                 l++;
//             } else {
//                 r++;
//             }
//             if (l == r) {
//                 max = Math.max(max, l * 2);
//             } else if (l > r) {
//                 l = r = 0;
//             }
//         }
//         return max;
//     }
// }
// @lc code=end

