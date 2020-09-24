import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=131 lang=java
 *
 * [131] 分割回文串
 *
 * https://leetcode-cn.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (63.91%)
 * Likes:    166
 * Dislikes: 0
 * Total Accepted:    14.3K
 * Total Submissions: 22.4K
 * Testcase Example:  '"aab"'
 *
 * 给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
 * 
 * 返回 s 所有可能的分割方案。
 * 
 * 示例:
 * 
 * 输入: "aab"
 * 输出:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 * 
 */

// @lc code=start
class Solution {
    private boolean isPalidrome(String sb) {
        int left = 0;
        int right = sb.length() - 1;
        while (left < right) {
            if (sb.charAt(left) != sb.charAt(right))
                return false;
            left++;
            right--;
        }
        return true;

    }

    private void backtrack(List<List<String>> res, String s, ArrayList<String> tmp) {
        if (s == null || s.length() == 0)
            res.add(tmp);
        for (int i = 1; i <= s.length(); i++) {
            if (isPalidrome(s.substring(0, i))) {
                tmp.add(s.substring(0, i));
                backtrack(res, s.substring(i, s.length()), tmp);
                tmp.remove(tmp.size() - 1);
            }
        }
    }

    public List<List<String>> partition(String s) {
        List<List<String>> res = new ArrayList<>();
        backtrack(res, s, new ArrayList<>());
        return res;
    }
}
// @lc code=end
