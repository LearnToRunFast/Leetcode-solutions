/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (53.47%)
 * Likes:    78
 * Dislikes: 0
 * Total Accepted:    13.7K
 * Total Submissions: 25.7K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
 *
 * 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
 *
 * 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
 *
 * 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
 *
 * 现在给定3个参数 — start, end,
 * bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回
 * -1。
 *
 * 注意：
 *
 *
 * 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
 * 如果一个起始基因序列需要多次变化，那么它每一次变化之后的基因序列都必须是合法的。
 * 假定起始基因序列与目标基因序列是不一样的。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * 返回值: 1
 *
 *
 * 示例 2：
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * 返回值: 2
 *
 *
 * 示例 3：
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * 返回值: 3
 *
 *
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {

	// bank map
	bankMap := map[string]bool{}
	for _, str := range bank {
		bankMap[str] = true
	}
	if len(bank) == 0 || !bankMap[end] {
		return -1
	}
	// replacement
	replacement := [4]byte{'A', 'C', 'G', 'T'}

	// two way bfs
	heads := map[string]bool{start: true}
	tails := map[string]bool{end: true}
	temp := map[string]bool{}

	ans := 0
	for len(heads) > 0 && len(tails) > 0 {
		// head is always the smaller one
		if len(heads) > len(tails) {
			heads, tails = tails, heads
		}
		ans += 1
		for str := range heads {
			head := []byte(str)
			for i, c := range head {
				old := c
				for _, rep := range replacement {
					head[i] = rep
					newStr := string(head)
					if tails[newStr] {
						return ans
					}
					if bankMap[newStr] {
						bankMap[newStr] = false
						temp[newStr] = true
					}
				}
				head[i] = old
			}
		}

		heads = temp
		temp = map[string]bool{}
	}
	return -1
}

// @lc code=end

