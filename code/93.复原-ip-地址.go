/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 *
 * https://leetcode-cn.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (51.70%)
 * Likes:    519
 * Dislikes: 0
 * Total Accepted:    102.7K
 * Total Submissions: 198.3K
 * Testcase Example:  '"25525511135"'
 *
 * 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
 *
 * 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
 *
 * 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312"
 * 和 "192.168@1.1" 是 无效 IP 地址。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "25525511135"
 * 输出：["255.255.11.135","255.255.111.35"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "0000"
 * 输出：["0.0.0.0"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "1111"
 * 输出：["1.1.1.1"]
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "010010"
 * 输出：["0.10.0.10","0.100.1.0"]
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "101023"
 * 输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * s 仅由数字组成
 *
 *
 */
// @lc code=start
const MAX_SEG_LEN = 4

func isValidSeg(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	n, err := strconv.Atoi(s)
	return err == nil && n >= 0 && n <= 255
}
func dfs(s string, start int, seg []string, res *[]string) {
	if len(seg) == MAX_SEG_LEN || start == len(s) {
		if len(seg) == MAX_SEG_LEN && start == len(s) {
			*res = append(*res, strings.Join(seg, "."))
		}
		return
	}
	for i := start; i < len(s); i++ {
		newSeg := s[start : i+1]
		if isValidSeg(newSeg) && len(seg) < 4 {
			dfs(s, i+1, append(seg, newSeg), res)
		} else {
			break
		}
	}
}
func restoreIpAddresses(s string) []string {
	var res []string
	var seg []string
	dfs(s, 0, seg, &res)
	return res
}

// @lc code=end
