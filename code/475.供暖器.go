/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 *
 * https://leetcode-cn.com/problems/heaters/description/
 *
 * algorithms
 * Medium (32.64%)
 * Likes:    205
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 57.9K
 * Testcase Example:  '[1,2,3]\n[2]'
 *
 * 冬季已经来临。 你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
 *
 * 在加热器的加热半径范围内的每个房屋都可以获得供暖。
 *
 * 现在，给出位于一条水平线上的房屋 houses 和供暖器 heaters 的位置，请你找出并返回可以覆盖所有房屋的最小加热半径。
 *
 * 说明：所有供暖器都遵循你的半径标准，加热的半径也一样。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: houses = [1,2,3], heaters = [2]
 * 输出: 1
 * 解释: 仅在位置2上有一个供暖器。如果我们将加热半径设为1，那么所有房屋就都能得到供暖。
 *
 *
 * 示例 2:
 *
 *
 * 输入: houses = [1,2,3,4], heaters = [1,4]
 * 输出: 1
 * 解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
 *
 *
 * 示例 3：
 *
 *
 * 输入：houses = [1,5], heaters = [2]
 * 输出：3
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 *
 *
 */

// @lc code=start
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findRadius(houses []int, heaters []int) int {
	sort.Slice(houses, func(i, j int) bool {
		return houses[i] < houses[j]
	})
	sort.Slice(heaters, func(i, j int) bool {
		return heaters[i] < heaters[j]
	})
	radius := 0
	i := 0
	for _, house := range houses {
		// find the closest right heater
		for i < len(heaters) && heaters[i] < house {
			i++
		}

		if i == 0 {
			radius = max(radius, heaters[i]-house)
		} else if i == len(heaters) {
			return max(radius, houses[len(houses)-1]-heaters[len(heaters)-1])
		} else {
			radius = max(radius, min(heaters[i]-house, house-heaters[i-1]))
		}
	}
	return radius
}

// @lc code=end

