import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=301 lang=java
 *
 * [301] 删除无效的括号
 *
 * https://leetcode-cn.com/problems/remove-invalid-parentheses/description/
 *
 * algorithms
 * Hard (47.92%)
 * Likes:    295
 * Dislikes: 0
 * Total Accepted:    16.3K
 * Total Submissions: 33.4K
 * Testcase Example:  '"()())()"'
 *
 * 删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。
 * 
 * 说明: 输入可能包含了除 ( 和 ) 以外的字符。
 * 
 * 示例 1:
 * 
 * 输入: "()())()"
 * 输出: ["()()()", "(())()"]
 * 
 * 
 * 示例 2:
 * 
 * 输入: "(a)())()"
 * 输出: ["(a)()()", "(a())()"]
 * 
 * 
 * 示例 3:
 * 
 * 输入: ")("
 * 输出: [""]
 * 
 */

// @lc code=start
class Solution {

    private void backTrack(String s, Set<String> res, StringBuilder sb, 
        int currIdx, int open, int close, int extraOpen, int extraClose) {

            if (currIdx == s.length()) {
                if (extraOpen == 0 && extraClose == 0) {
                    res.add(sb.toString());
                }
                return;
            }

            if (s.charAt(currIdx) == '(' && extraOpen > 0) {
                backTrack(s, res, sb, currIdx + 1, open, close, extraOpen - 1, extraClose);
            }

            if (s.charAt(currIdx) == ')' && extraClose > 0) {
                backTrack(s, res, sb, currIdx + 1, open, close, extraOpen, extraClose - 1);
            }
            
            sb.append(s.charAt(currIdx));
            
            if (s.charAt(currIdx) != '(' && s.charAt(currIdx) != ')') {
                backTrack(s, res, sb, currIdx + 1, open, close, extraOpen, extraClose);
            } else if (s.charAt(currIdx) == '(') {
                backTrack(s, res, sb, currIdx + 1, open + 1, close, extraOpen, extraClose);
            } else if (open > close) {
                backTrack(s, res, sb, currIdx + 1, open, close + 1, extraOpen, extraClose);
            }
            sb.deleteCharAt(sb.length() - 1);

    }
    public List<String> removeInvalidParentheses(String s) {
        int extraOpen = 0;
        int extraClose = 0;
        for (char c : s.toCharArray()) {
            if (c == '(') {
                extraOpen++;
                continue;
            }

            if (c == ')') {
                if (extraOpen > 0) {
                    extraOpen--;
                } else {
                    extraClose++;
                }
            }
        }
        Set<String> res = new HashSet<>();
        int currIdx = 0, open = 0, close = 0;
        backTrack(s, res, new StringBuilder(), currIdx, open, close, extraOpen, extraClose);
        return new ArrayList(res);
    }
}
// @lc code=end

