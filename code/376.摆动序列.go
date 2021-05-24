/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 *
 * https://leetcode-cn.com/problems/wiggle-subsequence/description/
 *
 * algorithms
 * Medium (45.37%)
 * Likes:    446
 * Dislikes: 0
 * Total Accepted:    60.3K
 * Total Submissions: 132.6K
 * Testcase Example:  '[1,7,4,9,2,5]'
 *
 * 如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为 摆动序列
 * 。第一个差（如果存在的话）可能是正数或负数。仅有一个元素或者含两个不等元素的序列也视作摆动序列。
 *
 *
 *
 * 例如， [1, 7, 4, 9, 2, 5] 是一个 摆动序列 ，因为差值 (6, -3, 5, -7, 3) 是正负交替出现的。
 *
 * 相反，[1, 4, 7, 2, 5] 和 [1, 7, 4, 5, 5]
 * 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。
 *
 *
 * 子序列 可以通过从原始序列中删除一些（也可以不删除）元素来获得，剩下的元素保持其原始顺序。
 *
 * 给你一个整数数组 nums ，返回 nums 中作为 摆动序列 的 最长子序列的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,7,4,9,2,5]
 * 输出：6
 * 解释：整个序列均为摆动序列，各元素之间的差值为 (6, -3, 5, -7, 3) 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,17,5,10,13,15,10,5,16,8]
 * 输出：7
 * 解释：这个序列包含几个长度为 7 摆动序列。
 * 其中一个是 [1, 17, 10, 13, 10, 16, 8] ，各元素之间的差值为 (16, -7, 3, -3, 6, -8) 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,2,3,4,5,6,7,8,9]
 * 输出：2
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
 *
 *
 * 进阶：你能否用 O(n) 时间复杂度完成此题?
 *
 */

// @lc code=start
// func wiggleMaxLength(nums []int) int {
// 	n := len(nums)
// 	dp := make([]int, n)
// 	signs := make([]int, n)
// 	max := 0
// 	for i := range nums {
// 		val := 1
// 		for j := i - 1; j >= 0; j-- {
// 			diff := nums[i] - nums[j]
// 			if diff != 0 {
// 				if signs[j] == 0 || signs[j] > 0 && diff < 0 || signs[j] < 0 && diff > 0 {
// 					val += dp[j]
// 					if diff < 0 {
// 						signs[i] = -1
// 					} else {
// 						signs[i] = 1
// 					}
// 					break
// 				}
// 			}
// 		}
// 		if max < val {
// 			max = val
// 		}
// 		dp[i] = val
// 	}
// 	return max
// }
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	inc, dec := 1, 1
	for i := 1; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 {
			inc = dec + 1
		} else if diff < 0 {
			dec = inc + 1
		}
	}
	ans := inc
	if ans < dec {
		ans = dec
	}
	return ans
}

// @lc code=end

