/*
 * @lc app=leetcode.cn id=76 lang=java
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode-cn.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (39.57%)
 * Likes:    827
 * Dislikes: 0
 * Total Accepted:    88.4K
 * Total Submissions: 223.3K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
 * 
 * 注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 和 t 由英文字母组成
 * 
 * 
 * 
 * 进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
 */

// @lc code=start
class Solution {
    public String minWindow(String s, String t) {
        int sLen = s.length(), tLen = t.length();
        if (sLen < tLen) return "";

        //ascii table size
        int[] window = new int[128]; // store window string
        int[] need = new int[128]; // store t string

        for (int i = 0; i < tLen; i++) {
            need[t.charAt(i)]++;
        }

        int l = 0, r = 0;
        int start = 0, len = sLen + 1;
        int needCount = 0;
        // [l, r)
        while (r < sLen) {
            char c = s.charAt(r);
            r++;

            if (need[c] > 0) {
                window[c]++;
                if (window[c] <= need[c]) {
                    needCount++;
                }
            }
            // move l until the window is no longer valid window for string "t"
            while (needCount == tLen) {
                if (len > r - l) {
                    start = l;
                    len = r - l;
                }

                // start moving index l
                char d = s.charAt(l);
                l++;
                
                if (need[d] > 0) {
                    if (window[d] <= need[d]) {
                        needCount--;
                    }
                    window[d]--;
                }
            }
        }
        return len == sLen + 1 ? "" : s.substring(start, start + len);
    }
}
// @lc code=end

