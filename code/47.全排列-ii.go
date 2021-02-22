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

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	ans := [][]int{}

	var backtrack func(int)

	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append(make([]int, 0, len(perm)), perm...))
			return
		}
		for i, v := range nums {
			// only take the first duplidated number that havent been visited before.
			if vis[i] || (i > 0 && v == nums[i-1] && !vis[i-1]) {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1] // remove last element
		}
	}

	backtrack(0)
	return ans
}

// @lc code=end

