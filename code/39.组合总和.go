/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.65%)
 * Likes:    1480
 * Dislikes: 0
 * Total Accepted:    309.9K
 * Total Submissions: 426.5K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target
 * 的唯一组合。
 *
 * candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
 *
 * 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: candidates = [2,3,6,7], target = 7
 * 输出: [[7],[2,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 *
 * 示例 3：
 *
 *
 * 输入: candidates = [2], target = 1
 * 输出: []
 *
 *
 * 示例 4：
 *
 *
 * 输入: candidates = [1], target = 1
 * 输出: [[1]]
 *
 *
 * 示例 5：
 *
 *
 * 输入: candidates = [1], target = 2
 * 输出: [[1,1]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * candidate 中的每个元素都是独一无二的。
 * 1
 *
 *
 */

// @lc code=start
func dfs(ans *[][]int, path, candidates []int, target, startIndex int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if target-candidates[i] < 0 {
			break
		}
		dfs(ans, append(path, candidates[i]), candidates, target-candidates[i], i)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) == 0 {
		return ans
	}
	sort.Ints(candidates)
	startIndex := 0
	dfs(&ans, []int{}, candidates, target, startIndex)
	return ans
}

// @lc code=end

