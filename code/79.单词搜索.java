/*
 * @lc app=leetcode.cn id=79 lang=java
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (43.74%)
 * Likes:    684
 * Dislikes: 0
 * Total Accepted:    120.6K
 * Total Submissions: 275.5K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 * 
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 * 
 * 
 * 
 * 示例:
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * 给定 word = "ABCCED", 返回 true
 * 给定 word = "SEE", 返回 true
 * 给定 word = "ABCB", 返回 false
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * board 和 word 中只包含大写和小写英文字母。
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 * 
 * 
 */

// @lc code=start
class Solution {

    private boolean dfs(int i, int j, int curr, char[][]board, char[] words) {
        if (i < 0 || i >= board.length || j < 0 || j >= board[0].length || board[i][j] != words[curr]) {
            return false;
        }

        if (curr == words.length - 1) {
            return true;
        }

        char c = board[i][j];
        board[i][j] = '-';
        boolean res = dfs(i - 1, j, curr + 1, board, words) || 
                      dfs(i, j - 1, curr + 1, board, words) || 
                      dfs(i + 1, j, curr + 1, board, words) || 
                      dfs(i, j + 1, curr + 1, board, words);

        board[i][j] = c;
        return res;

    }
    public boolean exist(char[][] board, String word) {
        char[] words = word.toCharArray();

        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (dfs(i, j, 0, board, words) ) {
                    return true;
                }
            }
        }
        return false;
    }   
}
// @lc code=end

