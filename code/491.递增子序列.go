/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 *
 * https://leetcode-cn.com/problems/increasing-subsequences/description/
 *
 * algorithms
 * Medium (54.92%)
 * Likes:    313
 * Dislikes: 0
 * Total Accepted:    42.3K
 * Total Submissions: 77.3K
 * Testcase Example:  '[4,6,7,7]'
 *
 * 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
 *
 * 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,6,7,7]
 * 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,4,3,2,1]
 * 输出：[[4,4]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 15
 * -100 <= nums[i] <= 100
 *
 *
 */

// @lc code=start
func dfs(nums []int, startIndex int, path []int, ans *[][]int) {
	if len(path) >= 2 {
		*ans = append(*ans, append([]int{}, path...))
	}
	visited := make([]int, 201)
	for i := startIndex; i < len(nums); i++ {
		// the range is from -100 to 100
		if visited[nums[i]+100] == 1 {
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			visited[nums[i]+100] = 1
			dfs(nums, i+1, append(path, nums[i]), ans)
		}

	}
}
func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	startIndex := 0
	dfs(nums, startIndex, path, &ans)
	return ans
}

// @lc code=end

