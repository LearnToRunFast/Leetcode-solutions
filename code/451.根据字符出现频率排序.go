/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 *
 * https://leetcode-cn.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (70.86%)
 * Likes:    320
 * Dislikes: 0
 * Total Accepted:    74.6K
 * Total Submissions: 105.1K
 * Testcase Example:  '"tree"'
 *
 * 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "tree"
 *
 * 输出:
 * "eert"
 *
 * 解释:
 * 'e'出现两次，'r'和't'都只出现一次。
 * 因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * "cccaaa"
 *
 * 输出:
 * "cccaaa"
 *
 * 解释:
 * 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
 * 注意"cacaca"是不正确的，因为相同的字母必须放在一起。
 *
 *
 * 示例 3:
 *
 *
 * 输入:
 * "Aabb"
 *
 * 输出:
 * "bbAa"
 *
 * 解释:
 * 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
 * 注意'A'和'a'被认为是两种不同的字符。
 *
 *
 */

// @lc code=start
func frequencySort(s string) string {
	dict := make([][2]int, 128)
	n := len(s)
	for i := range s {
		dict[s[i]][0]++
		dict[s[i]][1] = int(s[i])
	}

	ans := make([]byte, n)
	sort.Slice(dict, func(i, j int) bool { return dict[i][0] > dict[j][0] })
	curr := 0
	for i := range dict {
		if dict[i][0] == 0 {
			break
		}
		for j := 0; j < dict[i][0]; j++ {
			ans[curr] = byte(dict[i][1])
			curr++
		}
	}
	return string(ans)
}

// @lc code=end

