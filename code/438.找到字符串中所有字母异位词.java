import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=438 lang=java
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Medium (48.42%)
 * Likes:    441
 * Dislikes: 0
 * Total Accepted:    50.1K
 * Total Submissions: 103.3K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 * 
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 * 
 * 说明：
 * 
 * 
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入:
 * s: "cbaebabacd" p: "abc"
 * 
 * 输出:
 * [0, 6]
 * 
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入:
 * s: "abab" p: "ab"
 * 
 * 输出:
 * [0, 1, 2]
 * 
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 
 * 
 */

// @lc code=start
// class Solution {
//     public List<Integer> findAnagrams(String s, String p) {
//         if (p.length() > s.length()) return new ArrayList<Integer>();
//         int[] need = new int[128];
//         int[] window = new int[128];
//         for (char c : p.toCharArray()) {
//             need[c]++;
//         }
//         int l = 0, r = 0;
//         int needCount = 0;
//         List<Integer> ans = new ArrayList<>();
//         while (r < s.length()) {
//             char c = s.charAt(r);
//             r++;

//             if (need[c] > 0) {
//                 window[c]++;
//                 if (window[c] <= need[c]) {
//                     needCount++;
//                 }
//             }
//             while (r - l == p.length()) {
//                 if (needCount == p.length()) {
//                     ans.add(l);
//                 }
//                 char d = s.charAt(l);
//                 l++;

//                 if (need[d] > 0) {
//                     if (window[d] <= need[d]) {
//                         needCount--;
//                     }
//                     window[d]--;
//                 }
//             }
//         }
//         return ans;
//     }
// }
class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> ans = new ArrayList<>();
        if (p.length() > s.length()) return ans;
        int[] need = new int[128];
        int[] window = new int[128];
        for (char c : p.toCharArray()) {
            need[c]++;
        }
        int l = 0, r = 0;

        while (r < s.length()) {
            char c = s.charAt(r);
            r++;

            window[c]++;

            while (window[c] > need[c]) {
                char d = s.charAt(l);
                l++;
                window[d]--;
            }
            if (r - l == p.length()) {
                ans.add(l);
            }
        }
        return ans;
    }
}
// @lc code=end

