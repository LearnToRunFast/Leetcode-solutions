/*
 * @lc app=leetcode.cn id=22 lang=java
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (76.06%)
 * Likes:    1295
 * Dislikes: 0
 * Total Accepted:    176.7K
 * Total Submissions: 232K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例：
 * 
 * 输入：n = 3
 * 输出：[
 * ⁠      "((()))",
 * ⁠      "(()())",
 * ⁠      "(())()",
 * ⁠      "()(())",
 * ⁠      "()()()"
 * ⁠    ]
 * 
 * 
 */

// @lc code=start
class Solution {
    public List<String> generateParenthesis(int n) {
        List<String> list = new ArrayList<>();
        dfs(list, "", 0, 0, n);
        return list;
    }
    
    private void backtrack(List<String> list, String str, int open, int close, int max) {
        if (str.length() == max * 2) {
            list.add(str);
            return;
        }

        if (open < close) {
            return;
        }


        if (open < max) {
            dfs(list, str +"(", open + 1, close, max);
        }
        if (close < max) {
            dfs(list, str +")", open , close + 1, max);
        }


    }
}
// @lc code=end

