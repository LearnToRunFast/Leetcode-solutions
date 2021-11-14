/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 *
 * https://leetcode-cn.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (61.77%)
 * Likes:    404
 * Dislikes: 0
 * Total Accepted:    69.1K
 * Total Submissions: 111.7K
 * Testcase Example:  '[1,2,2]'
 *
 * 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 *
 * 说明：解集不能包含重复的子集。
 *
 * 示例:
 *
 * 输入: [1,2,2]
 * 输出:
 * [
 * ⁠ [2],
 * ⁠ [1],
 * ⁠ [1,2,2],
 * ⁠ [2,2],
 * ⁠ [1,2],
 * ⁠ []
 * ]
 *
 */
// @lc code=start
func dfs(nums []int, startIndex int, perm []int, res *[][]int) {
	*res = append(*res, append([]int{}, perm...))
	if startIndex == len(nums) {
		return
	}
	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}
		dfs(nums, i+1, append(perm, nums[i]), res)
	}
}
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	startIndex := 0
	sort.Ints(nums)
	dfs(nums, startIndex, []int{}, &res)
	return res
}

// @lc code=end
