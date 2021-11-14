/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (61.96%)
 * Likes:    731
 * Dislikes: 0
 * Total Accepted:    212.4K
 * Total Submissions: 342.9K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 注意：解集不能包含重复的组合。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 * 示例 2:
 *
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * 提示:
 *
 *
 * 1
 * 1
 * 1
 *
 *
 */

// @lc code=start
func dfs(ans *[][]int, candidates, path []int, target, startIndex int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		dfs(ans, candidates, append(path, candidates[i]), target-candidates[i], i+1)
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) == 0 || target == 0 {
		return ans
	}
	sort.Ints(candidates)
	dfs(&ans, candidates, []int{}, target, 0)
	return ans
}

// @lc code=end

