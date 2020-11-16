/*
 * @lc app=leetcode.cn id=647 lang=java
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (64.56%)
 * Likes:    436
 * Dislikes: 0
 * Total Accepted:    70.4K
 * Total Submissions: 109.1K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
 * 
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 * 
 * 
 * 示例 2：
 * 
 * 输入："aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的字符串长度不会超过 1000 。
 * 
 * 
 */

// @lc code=start
//DP
// class Solution {
//     public int countSubstrings(String s) {
//         int n = s.length();
//         boolean[] dp = new boolean[n];

//         int count = 0;
//         for (int j = 0; j < n; j++) {
//             for (int i = 0; i <= j; i++) {
//                 if (s.charAt(i) == s.charAt(j) &&
//                     (j - i < 2 || dp[i + 1])) {
//                         count++;
//                         dp[i] = true;
                    
//                 } else {
//                     dp[i] = false;
//                 }
//             } 
//         }
//         return count;
        
//     }
// }

class Solution {
    public int countSubstrings(String s) {
        int n = s.length();
        int count = 0;
        // 2 * n - 1 center points
        for (int center = 0; center < 2 * n - 1; center++) {
            int l = center / 2;
            int r = l + center % 2;

            while (l >= 0 && r < n && s.charAt(l) == s.charAt(r)) {
                count++;
                l--;
                r++;
            }
        }
        return count;
        
    }
}
// @lc code=end

