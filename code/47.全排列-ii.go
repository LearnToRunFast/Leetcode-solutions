/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (62.50%)
 * Likes:    596
 * Dislikes: 0
 * Total Accepted:    138.7K
 * Total Submissions: 221.1K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 *
 *
 */

// @lc code=start
func dfs(nums []int, path []int, res *[][]int, visited []bool) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		dfs(nums, append(path, nums[i]), res, visited)
		visited[i] = false
	}

}
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	perm := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums, perm, &res, visited)
	return res
}

// @lc code=end

