/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 *
 * https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/description/
 *
 * algorithms
 * Medium (47.86%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    22.1K
 * Total Submissions: 46.1K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n' +
  '[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
 *
 * 实现词典类 WordDictionary ：
 *
 *
 * WordDictionary() 初始化词典对象
 * void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
 * bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些
 * '.' ，每个 . 都可以表示任何一个字母。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 *
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * 输出：
 * [null,null,null,null,false,true,true,true]
 *
 * 解释：
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // return False
 * wordDictionary.search("bad"); // return True
 * wordDictionary.search(".ad"); // return True
 * wordDictionary.search("b.."); // return True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * addWord 中的 word 由小写英文字母组成
 * search 中的 word 由 '.' 或小写英文字母组成
 * 最多调用 50000 次 addWord 和 search
 *
 *
*/
package main

// @lc code=start
// type WordDictionary struct {
// 	end   bool
// 	nexts map[byte]*WordDictionary
// }

// /** Initialize your data structure here. */
// func Constructor() WordDictionary {
// 	return WordDictionary{nexts: map[byte]*WordDictionary{}}
// }

// func (this *WordDictionary) AddWord(word string) {
// 	root := this
// 	for i := range word {
// 		c := word[i]
// 		if _, ok := root.nexts[c]; !ok {
// 			root.nexts[c] = &WordDictionary{nexts: map[byte]*WordDictionary{}}
// 		}
// 		root = root.nexts[c]
// 	}
// 	root.end = true
// }
// func (this *WordDictionary) backTrack(word *string, curr int) bool {
// 	if curr == len(*word) {
// 		return this.end
// 	}
// 	c := (*word)[curr]
// 	if c == '.' {
// 		for _, v := range this.nexts {
// 			if v.backTrack(word, curr+1) {
// 				return true
// 			}
// 		}
// 	} else {

// 		next, ok := this.nexts[c]
// 		if !ok {
// 			return false
// 		}
// 		return next.backTrack(word, curr+1)
// 	}
// 	return false
// }
// func (this *WordDictionary) Search(word string) bool {
// 	return this.backTrack(&word, 0)
// }
type WordDictionary struct {
	end   bool
	nexts map[byte]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{nexts: map[byte]*WordDictionary{}}
}

func (this *WordDictionary) AddWord(word string) {
	root := this
	for i := range word {
		c := word[i]
		if _, ok := root.nexts[c]; !ok {
			root.nexts[c] = &WordDictionary{nexts: map[byte]*WordDictionary{}}
		}
		root = root.nexts[c]
	}
	root.end = true
}
func (this *WordDictionary) backTrack(word string) bool {
	if word == "" {
		return this.end
	}
	c := word[0]
	if c == '.' {
		for _, v := range this.nexts {
			if v.backTrack(word[1:]) {
				return true
			}
		}
	} else {

		next, ok := this.nexts[c]
		if !ok {
			return false
		}
		return next.backTrack(word[1:])
	}
	return false
}
func (this *WordDictionary) Search(word string) bool {
	return this.backTrack(word[:])
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
