/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (53.88%)
 * Likes:    255
 * Dislikes: 0
 * Total Accepted:    28.1K
 * Total Submissions: 52.2K
 * Testcase Example:  '1'
 *
 * 二进制手表顶部有 4 个 LED 代表 小时（0-11），底部的 6 个 LED 代表 分钟（0-59）。每个 LED 代表一个 0 或
 * 1，最低位在右侧。
 *
 *
 * 例如，下面的二进制手表读取 "3:25" 。
 *
 *
 *
 *
 * （图源：WikiMedia - Binary clock samui moon.jpg ，许可协议：Attribution-ShareAlike 3.0
 * Unported (CC BY-SA 3.0) ）
 *
 * 给你一个整数 turnedOn ，表示当前亮着的 LED 的数量，返回二进制手表可以表示的所有可能时间。你可以 按任意顺序 返回答案。
 *
 * 小时不会以零开头：
 *
 *
 * 例如，"01:00" 是无效的时间，正确的写法应该是 "1:00" 。
 *
 *
 * 分钟必须由两位数组成，可能会以零开头：
 *
 *
 * 例如，"10:2" 是无效的时间，正确的写法应该是 "10:02" 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：turnedOn = 1
 * 输出：["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：turnedOn = 9
 * 输出：[]
 *
 *
 *
 *
 * 解释：
 *
 *
 * 0
 *
 *
 */

// @lc code=start
func dfs(ans *[]string, num, h, m, start int) {

	if num == 0 {
		val := fmt.Sprintf("%v:%02d", h, m)
		*ans = append(*ans, val)
	}
	for i := start; i < 10; i++ {
		if i < 6 {
			val := m + 1<<i
			if val > 59 {
				continue
			}
			dfs(ans, num-1, h, val, i+1)
		} else {
			val := h + 1<<(i-6)
			if val > 11 {
				return
			}
			dfs(ans, num-1, val, m, i+1)
		}
	}
}
func readBinaryWatch(turnedOn int) []string {
	ans := []string{}
	dfs(&ans, turnedOn, 0, 0, 0)
	return ans
}

// @lc code=end

