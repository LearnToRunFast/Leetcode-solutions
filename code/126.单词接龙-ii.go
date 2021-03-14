/*
 * @lc app=leetcode.cn id=126 lang=golang
 *
 * [126] 单词接龙 II
 *
 * https://leetcode-cn.com/problems/word-ladder-ii/description/
 *
 * algorithms
 * Hard (39.17%)
 * Likes:    403
 * Dislikes: 0
 * Total Accepted:    29.6K
 * Total Submissions: 75.4K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord
 * 的最短转换序列。转换需遵循如下规则：
 *
 *
 * 每次转换只能改变一个字母。
 * 转换后得到的单词必须是字典中的单词。
 *
 *
 * 说明:
 *
 *
 * 如果不存在这样的转换序列，返回一个空列表。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 *
 *
 * 示例 1:
 *
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * 输出:
 * [
 * ⁠ ["hit","hot","dot","dog","cog"],
 * ["hit","hot","lot","log","cog"]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * 输出: []
 *
 * 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。
 *
 */

// @lc code=start
package main

func getNeighbours(word string, dict map[string]bool) []string {
	neighbour := []string{}
	chars := []rune(word)
	for i, v := range chars {
		tmp := v
		for c := 'a'; c <= 'z'; c++ {
			if c == tmp {
				continue
			}
			chars[i] = c
			newStr := string(chars)
			if dict[newStr] {
				neighbour = append(neighbour, newStr)
			}
		}
		chars[i] = tmp
	}
	return neighbour
}
func getGraph(edges map[string][]string, dict map[string]bool, endWord string) map[string]int {
	dist := map[string]int{endWord: 0}
	queue := []string{endWord}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[0]
			queue = queue[1:]
			neighbours := getNeighbours(word, dict)
			for _, v := range neighbours {
				edges[v] = append(edges[v], word)
				if _, ok := dist[v]; !ok {
					dist[v] = dist[word] + 1
					queue = append(queue, v)
				}
			}
		}
	}
	return dist
}
func backtracking(beginWord, endword string, paths *[][]string, path []string, dist map[string]int, edges map[string][]string) {
	if beginWord == endword {
		path = append(path, endword)
		(*paths) = append((*paths), append([]string{}, path...))
		path = path[:len(path)-1]
		return
	}
	for _, v := range edges[beginWord] {
		if dist[beginWord] == dist[v]+1 {
			path = append(path, beginWord)
			backtracking(v, endword, paths, path, dist, edges)
			path = path[:len(path)-1]
		}
	}
}
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	paths := [][]string{}
	dict := map[string]bool{}
	wordList = append(wordList, beginWord)
	for _, v := range wordList {
		dict[v] = true
	}
	if _, ok := dict[endWord]; !ok {
		return paths
	}
	edges := map[string][]string{}
	dist := getGraph(edges, dict, endWord)
	backtracking(beginWord, endWord, &paths, []string{}, dist, edges)
	return paths
}

// @lc code=end
