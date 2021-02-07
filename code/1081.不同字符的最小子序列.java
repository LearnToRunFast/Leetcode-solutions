/*
 * @lc app=leetcode.cn id=1081 lang=java
 *
 * [1081] 不同字符的最小子序列
 *
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters/description/
 *
 * algorithms
 * Medium (56.76%)
 * Likes:    81
 * Dislikes: 0
 * Total Accepted:    9.6K
 * Total Submissions: 16.9K
 * Testcase Example:  '"bcabc"'
 *
 * 返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
 * 
 * 注意：该题与 316 https://leetcode.com/problems/remove-duplicate-letters/ 相同
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "bcabc"
 * 输出："abc"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "cbacdcbc"
 * 输出："acdb"
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 由小写英文字母组成
 * 
 * 
 */

// @lc code=start
class Solution {
    public String smallestSubsequence(String s) {
        int[] dict = new int[26];
        boolean[] seen = new boolean[26];
        StringBuilder sb = new StringBuilder();
        char[] ss = s.toCharArray();
        for (char c : ss) {
            dict[c - 'a']++;
        }

        for (char c : ss) {
            if (seen[c - 'a']) {
                dict[c - 'a']--;
                continue;
            }

            while (sb.length() > 0 && sb.charAt(sb.length() - 1) > c && dict[sb.charAt(sb.length() - 1) - 'a'] > 1) {
                int lastChar = sb.charAt(sb.length() - 1) - 'a';
                sb.deleteCharAt(sb.length() - 1);
                seen[lastChar] = false;
                dict[lastChar]--;
            }
            sb.append(c);
            seen[c - 'a'] = true;
        }
        return sb.toString();
    }
}
// @lc code=end

