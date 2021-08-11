/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 *
 * https://leetcode-cn.com/problems/matchsticks-to-square/description/
 *
 * algorithms
 * Medium (41.63%)
 * Likes:    193
 * Dislikes: 0
 * Total Accepted:    18.8K
 * Total Submissions: 45K
 * Testcase Example:  '[1,1,2,2,2]'
 *
 *
 * 还记得童话《卖火柴的小女孩》吗？现在，你知道小女孩有多少根火柴，请找出一种能使用所有火柴拼成一个正方形的方法。不能折断火柴，可以把火柴连接起来，并且每根火柴都要用到。
 *
 * 输入为小女孩拥有火柴的数目，每根火柴用其长度表示。输出即为是否能用所有的火柴拼成正方形。
 *
 * 示例 1:
 *
 *
 * 输入: [1,1,2,2,2]
 * 输出: true
 *
 * 解释: 能拼成一个边长为2的正方形，每边两根火柴。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,3,3,3,4]
 * 输出: false
 *
 * 解释: 不能用所有火柴拼成一个正方形。
 *
 *
 * 注意:
 *
 *
 * 给定的火柴长度和在 0 到 10^9之间。
 * 火柴数组的长度不超过15。
 *
 *
 */

// @lc code=start
func dfs(nums []int, curr, target int, sides []int) bool {
	if curr == len(nums) {
		for i := 1; i < len(sides); i++ {
			if sides[0] != sides[i] {
				return false
			}
		}
		return true
	}
	for i, side := range sides {
		if side+nums[curr] > target || (i > 0 && side == sides[i-1]) || (curr == 0 && i > 0) {
			continue
		}

		sides[i] += nums[curr]
		if dfs(nums, curr+1, target, sides) {
			return true
		}
		sides[i] -= nums[curr]
	}
	return false
}
func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}
	var sum = 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	return dfs(matchsticks, 0, sum/4, make([]int, 4))
}

// @lc code=end

