/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (49.90%)
 * Likes:    345
 * Dislikes: 0
 * Total Accepted:    93.6K
 * Total Submissions: 187.6K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
 *
 *
 * 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：s = "egg", t = "add"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "foo", t = "bar"
 * 输出：false
 *
 * 示例 3：
 *
 *
 * 输入：s = "paper", t = "title"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 可以假设 s 和 t 长度相同。
 *
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	lookup := map[byte]byte{}
	seen := map[byte]bool{}
	for i := range s {
		v := s[i]
		u := t[i]
		mappedVal, ok := lookup[v]
		if !ok {
			if seen[u] {
				return false
			}
			lookup[v] = u
			seen[u] = true
			continue
		}
		if mappedVal != t[i] {
			return false
		}
	}
	return true
}

// func isIsomorphic(s string, t string) bool {
// 	sMap, tMap := map[uint8]int{}, map[uint8]int{}
// 	for i := range s {
// 		var s_v uint8 = s[i]
// 		var t_v uint8 = t[i]
// 		if sMap[s_v] != tMap[t_v] {
// 			return false
// 		}
// 		val := i + 1
// 		sMap[s_v] = val
// 		tMap[t_v] = val
// 	}
// 	return true
// }

// @lc code=end

