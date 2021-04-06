/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 *
 * https://leetcode-cn.com/problems/word-search-ii/description/
 *
 * algorithms
 * Hard (45.67%)
 * Likes:    350
 * Dislikes: 0
 * Total Accepted:    31.3K
 * Total Submissions: 68.6K
 * Testcase Example:  '[["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]]\n' +
  '["oath","pea","eat","rain"]'
 *
 * 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
 *
 * 单词必须按照字母顺序，通过 相邻的单元格
 * 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：board =
 * [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
 * words = ["oath","pea","eat","rain"]
 * 输出：["eat","oath"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == board.length
 * n == board[i].length
 * 1
 * board[i][j] 是一个小写英文字母
 * 1
 * 1
 * words[i] 由小写英文字母组成
 * words 中的所有字符串互不相同
 *
 *
*/
package main

// @lc code=start
type Node struct {
	next [26]*Node
	word string
}

func (n *Node) addWord(word string) {
	root := n
	for i := range word {
		pos := word[i] - 'a'
		if root.next[pos] == nil {
			root.next[pos] = &Node{}
		}
		root = root.next[pos]
	}
	root.word = word
}

// return the next root and whether there is a match
func (n *Node) getPrefix(c byte) *Node {
	pos := c - 'a'
	return n.next[pos]
}

func findWords(board [][]byte, words []string) []string {
	root := &Node{}
	for _, v := range words {
		root.addWord(v)
	}
	n, m := len(board), len(board[0])
	ans := []string{}

	var dfs func(curr *Node, i, j int, ans *[]string)
	dfs = func(curr *Node, i, j int, ans *[]string) {
		if i < 0 || i >= n || j < 0 || j >= m || board[i][j] == '.' {
			return
		}
		c := board[i][j]
		nextPtr := curr.getPrefix(c)
		if nextPtr == nil {
			return
		}

		if nextPtr.word != "" {
			*ans = append(*ans, nextPtr.word)
			nextPtr.word = ""
		}

		board[i][j] = '.'
		dfs(nextPtr, i+1, j, ans)
		dfs(nextPtr, i, j+1, ans)
		dfs(nextPtr, i-1, j, ans)
		dfs(nextPtr, i, j-1, ans)
		board[i][j] = c
	}
	for i, row := range board {
		for j, _ := range row {
			dfs(root, i, j, &ans)
		}
	}
	return ans

}

// @lc code=end
