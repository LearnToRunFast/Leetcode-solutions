/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 *
 * https://leetcode-cn.com/problems/largest-divisible-subset/description/
 *
 * algorithms
 * Medium (45.63%)
 * Likes:    347
 * Dislikes: 0
 * Total Accepted:    37.6K
 * Total Submissions: 82.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个由 无重复 正整数组成的集合 nums ，请你找出并返回其中最大的整除子集 answer ，子集中每一元素对 (answer[i],
 * answer[j]) 都应当满足：
 *
 * answer[i] % answer[j] == 0 ，或
 * answer[j] % answer[i] == 0
 *
 *
 * 如果存在多个有效解子集，返回其中任何一个均可。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[1,2]
 * 解释：[1,3] 也会被视为正确答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,4,8]
 * 输出：[1,2,4,8]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * nums 中的所有整数 互不相同
 *
 *
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	if n < 2 {
		return nums
	}

	dp := make([]int, n)
	tracks := make([]int, n)
	max := 1
	maxId := 0
	for i := range nums {
		// length
		l := 1
		prev := i
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && dp[j]+1 > l {
				l = dp[j] + 1
				prev = j
				if max < l {
					max = l
					maxId = i
				}
			}
		}
		dp[i] = l
		tracks[i] = prev
	}
	ans := make([]int, max)
	for i := 0; i < max; i++ {
		ans[i] = nums[maxId]
		maxId = tracks[maxId]
	}
	return ans

}

// @lc code=end

