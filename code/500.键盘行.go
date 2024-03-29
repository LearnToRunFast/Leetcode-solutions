/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 *
 * https://leetcode-cn.com/problems/keyboard-row/description/
 *
 * algorithms
 * Easy (70.77%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    29K
 * Total Submissions: 41.1K
 * Testcase Example:  '["Hello","Alaska","Dad","Peace"]'
 *
 * 给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
 *
 * 美式键盘 中：
 *
 *
 * 第一行由字符 "qwertyuiop" 组成。
 * 第二行由字符 "asdfghjkl" 组成。
 * 第三行由字符 "zxcvbnm" 组成。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：words = ["Hello","Alaska","Dad","Peace"]
 * 输出：["Alaska","Dad"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：words = ["omk"]
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：words = ["adsdf","sfd"]
 * 输出：["adsdf","sfd"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * words[i] 由英文字母（小写和大写字母）组成
 *
 *
 */

// @lc code=start
func findWords(words []string) []string {
	var r = make([]string, 0)
	var dict = []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}
	for _, w := range words {
		word := strings.ToLower(w)
		var flag = true
		for _, m := range dict {
			if strings.ContainsAny(m, word) {
				for _, c := range word {
					if !strings.ContainsRune(m, c) {
						flag = false
						break
					}
				}
			}
			if !flag {
				break
			}
		}
		if flag {
			r = append(r, w)
		}
	}
	return r
}

// @lc code=end

