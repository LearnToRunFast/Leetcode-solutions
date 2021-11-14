/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (80.17%)
 * Likes:    1383
 * Dislikes: 0
 * Total Accepted:    329.9K
 * Total Submissions: 411.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */

// @lc code=start
func dfs(nums []int, startIndex int, subSet []int, res *[][]int) {
	*res = append(*res, append([]int{}, subSet...))
	if startIndex == len(nums) {
		return
	}
	for i := startIndex; i < len(nums); i++ {
		dfs(nums, i+1, append(subSet, nums[i]), res)
	}
}
func subsets(nums []int) [][]int {
	var res [][]int
	var subSet []int
	var startIndex int
	sort.Ints(nums)
	dfs(nums, startIndex, subSet, &res)
	return res
}

// @lc code=end

