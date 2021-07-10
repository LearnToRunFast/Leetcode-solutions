/*
 * @lc app=leetcode.cn id=436 lang=golang
 *
 * [436] 寻找右区间
 *
 * https://leetcode-cn.com/problems/find-right-interval/description/
 *
 * algorithms
 * Medium (48.86%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    7.3K
 * Total Submissions: 15K
 * Testcase Example:  '[[1,2]]'
 *
 * 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
 *
 * 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
 *
 * 返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：intervals = [[1,2]]
 * 输出：[-1]
 * 解释：集合中只有一个区间，所以输出-1。
 *
 *
 * 示例 2：
 *
 *
 * 输入：intervals = [[3,4],[2,3],[1,2]]
 * 输出：[-1, 0, 1]
 * 解释：对于 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间[3,4]具有最小的“右”起点;
 * 对于 [1,2] ，区间[2,3]具有最小的“右”起点。
 *
 *
 * 示例 3：
 *
 *
 * 输入：intervals = [[1,4],[2,3],[3,4]]
 * 输出：[-1, 2, -1]
 * 解释：对于区间 [1,4] 和 [3,4] ，没有满足条件的“右侧”区间。
 * 对于 [2,3] ，区间 [3,4] 有最小的“右”起点。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * intervals[i].length == 2
 * -10^6 i i
 * 每个间隔的起点都 不相同
 *
 *
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([]int, n)
	for i := range intervals {
		found := sort.Search(n,
			func(j int) bool {
				return intervals[i][1] <= intervals[j][0]
			})
		if found == n {
			ans[intervals[i][2]] = -1
		} else {
			ans[intervals[i][2]] = intervals[found][2]
		}
	}
	return ans
}

// @lc code=end

