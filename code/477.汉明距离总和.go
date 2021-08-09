/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 *
 * https://leetcode-cn.com/problems/total-hamming-distance/description/
 *
 * algorithms
 * Medium (60.28%)
 * Likes:    227
 * Dislikes: 0
 * Total Accepted:    38.1K
 * Total Submissions: 63.2K
 * Testcase Example:  '[4,14,2]'
 *
 * 两个整数的 汉明距离 指的是这两个数字的二进制数对应位不同的数量。
 *
 * 给你一个整数数组 nums，请你计算并返回 nums 中任意两个数之间汉明距离的总和。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,14,2]
 * 输出：6
 * 解释：在二进制表示中，4 表示为 0100 ，14 表示为 1110 ，2表示为 0010 。（这样表示是为了体现后四位之间关系）
 * 所以答案为：
 * HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2
 * + 2 + 2 = 6
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,14,4]
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
 *
 *
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	n := len(nums)
	count := 0
	for i := 0; i < 32; i++ {
		ones := 0
		zeros := 0
		for _, num := range nums {
			ones += num >> i & 0x01
		}
		zeros += n - ones
		count += zeros * ones
	}
	return count
}

// @lc code=end

