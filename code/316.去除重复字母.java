import java.util.HashSet;
import java.util.LinkedList;
import java.util.Set;
import java.util.Stack;

/*
 * @lc app=leetcode.cn id=316 lang=java
 *
 * [316] 去除重复字母
 *
 * https://leetcode-cn.com/problems/remove-duplicate-letters/description/
 *
 * algorithms
 * Medium (47.64%)
 * Likes:    452
 * Dislikes: 0
 * Total Accepted:    47.8K
 * Total Submissions: 100.3K
 * Testcase Example:  '"bcabc"'
 *
 * 给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
 * 
 * 注意：该题与 1081
 * https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters
 * 相同
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
    public String removeDuplicateLetters(String s) {
        int[] dict = new int[26]; // count appear times
        boolean[] seen = new boolean[26]; // check is visited
        for (char c : s.toCharArray()) {
            dict[c - 'a']++;
        }
        StringBuilder sb = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (seen[c - 'a']) {
                dict[c - 'a']--;
                continue;
            }
            
            while (sb.length() > 0 && sb.charAt(sb.length() - 1) > c && dict[sb.charAt(sb.length() - 1) - 'a'] > 1) {
                int lastCharIdx = sb.charAt(sb.length() - 1) - 'a';
                sb.deleteCharAt(sb.length() - 1);
                seen[lastCharIdx] = false;
                dict[lastCharIdx]--;
            }
            seen[c - 'a'] = true;
            sb.append(c);
        }
        return sb.toString();
    }
}
// @lc code=end

