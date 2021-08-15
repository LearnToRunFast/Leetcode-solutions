/*
 * @lc app=leetcode.cn id=539 lang=golang
 *
 * [539] 最小时间差
 *
 * https://leetcode-cn.com/problems/minimum-time-difference/description/
 *
 * algorithms
 * Medium (59.11%)
 * Likes:    94
 * Dislikes: 0
 * Total Accepted:    13K
 * Total Submissions: 22K
 * Testcase Example:  '["23:59","00:00"]'
 *
 * 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：timePoints = ["23:59","00:00"]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：timePoints = ["00:00","23:59","00:00"]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2
 * timePoints[i] 格式为 "HH:MM"
 *
 *
 */

// @lc code=start
func toNum(s string) int {
	base := 1
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		sum += int(s[i]-'0') * base
		base *= 10
	}
	return sum
}
func timeToMin(timePoints string) int {
	var times = strings.Split(timePoints, ":")
	h := toNum(times[0])
	m := toNum(times[1])
	return h*60 + m
}
func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	var times = make([]int, n)
	end := 24 * 60
	for i, v := range timePoints {
		times[i] = timeToMin(v)
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})
	times = append(times, 0) // any value is ok
	times[n] = end + times[0]
	min := 1<<31 - 1

	for i := 1; i <= n; i++ {
		diff := times[i] - times[i-1]
		if min > diff {
			min = diff
		}
	}

	return min

}

// @lc code=end

