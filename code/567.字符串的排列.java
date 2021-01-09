/*
 * @lc app=leetcode.cn id=567 lang=java
 *
 * [567] 字符串的排列
 *
 * https://leetcode-cn.com/problems/permutation-in-string/description/
 *
 * algorithms
 * Medium (37.85%)
 * Likes:    216
 * Dislikes: 0
 * Total Accepted:    51.4K
 * Total Submissions: 134.6K
 * Testcase Example:  '"ab"\n"eidbaooo"'
 *
 * 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
 * 
 * 换句话说，第一个字符串的排列之一是第二个字符串的子串。
 * 
 * 示例1:
 * 
 * 
 * 输入: s1 = "ab" s2 = "eidbaooo"
 * 输出: True
 * 解释: s2 包含 s1 的排列之一 ("ba").
 * 
 * 
 * 
 * 
 * 示例2:
 * 
 * 
 * 输入: s1= "ab" s2 = "eidboaoo"
 * 输出: False
 * 
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 输入的字符串只包含小写字母
 * 两个字符串的长度都在 [1, 10,000] 之间
 * 
 * 
 */

// @lc code=start
class Solution {
    public boolean checkInclusion(String s1, String s2) {
        if (s1.length() > s2.length()) return false;
        int asciiSize = 128;
        int[] window = new int[asciiSize];
        int[] need = new int[asciiSize];

        for (char c : s1.toCharArray()) {
            need[c]++;
        }

        int l = 0, r = 0;
        int needCount = 0;
        while (r < s2.length()) {
            char c = s2.charAt(r);
            r++;

            if (need[c] > 0) {
                window[c]++;
                if (window[c] <= need[c]) {
                    needCount++;
                }
            }
            if (r - l == s1.length()) {
                if (needCount == s1.length()) {
                    return true;
                }
                char d = s2.charAt(l);
                l++;
                if (need[d] > 0) {
                    if (window[d] <= need[d]) {
                        needCount--;
                    }
                    window[d]--;
                }
            }
        }
        return false;
    }
}
// @lc code=end

