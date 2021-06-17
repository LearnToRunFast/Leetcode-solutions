/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 *
 * https://leetcode-cn.com/problems/split-array-largest-sum/description/
 *
 * algorithms
 * Hard (55.40%)
 * Likes:    487
 * Dislikes: 0
 * Total Accepted:    33.4K
 * Total Submissions: 60.3K
 * Testcase Example:  '[7,2,5,10,8]\n2'
 *
 * 给定一个非负整数数组 nums 和一个整数 m ，你需要将这个数组分成 m 个非空的连续子数组。
 *
 * 设计一个算法使得这 m 个子数组各自和的最大值最小。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [7,2,5,10,8], m = 2
 * 输出：18
 * 解释：
 * 一共有四种方法将 nums 分割为 2 个子数组。 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
 * 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4,5], m = 2
 * 输出：9
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,4,4], m = 3
 * 输出：4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * 1
 *
 *
 */

// @lc code=start
func check(nums []int, x, m int) bool {
	sum, count := 0, 1
	for _, num := range nums {
		if sum+num > x {
			count++
			sum = num
		} else {
			sum += num
		}
	}
	return count <= m
}
func splitArray(nums []int, m int) int {
	max, sum := 0, 0
	for _, num := range nums {
		if num > max {
			max = num
		}
		sum += num
	}
	l, r := max, sum
	for l < r {
		mid := l + (r-l)/2
		if check(nums, mid, m) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// @lc code=end

