/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] 太平洋大西洋水流问题
 *
 * https://leetcode-cn.com/problems/pacific-atlantic-water-flow/description/
 *
 * algorithms
 * Medium (46.47%)
 * Likes:    257
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 47.7K
 * Testcase Example:  '[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]'
 *
 * 给定一个 n x m 的非负整数矩阵来表示一片大陆上各个单元格的高度。“太平洋”处于大陆的左边界和上边界，而“大西洋”处于大陆的右边界和下边界。
 *
 * 规定水流只能按照上、下、左、右四个方向流动，且只能从高到低或者在同等高度上流动。
 *
 * 请找出那些水流既可以流动到“太平洋”，又能流动到“大西洋”的陆地单元的坐标。
 *
 *
 *
 * 提示：
 *
 *
 * 输出坐标的顺序不重要
 * n 和 m 都小于150
 *
 *
 *
 *
 * 示例：
 *
 *
 *
 *
 * 给定下面的 5x5 矩阵:
 *
 * ⁠ 太平洋 ~   ~   ~   ~   ~
 * ⁠      ~  1   2   2   3  (5) *
 * ⁠      ~  3   2   3  (4) (4) *
 * ⁠      ~  2   4  (5)  3   1  *
 * ⁠      ~ (6) (7)  1   4   5  *
 * ⁠      ~ (5)  1   1   2   4  *
 * ⁠         *   *   *   *   * 大西洋
 *
 * 返回:
 *
 * [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (上图中带括号的单元).
 *
 *
 *
 *
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	ans := [][]int{}
	n := len(heights)
	if n == 0 {
		return ans
	}
	m := len(heights[0])
	a := make([][]bool, n)
	p := make([][]bool, n)
	for i := range a {
		a[i] = make([]bool, m)
		p[i] = make([]bool, m)
	}

	var dfs func(i, j, prevH int, seen *[][]bool)
	dfs = func(i, j, prevH int, seen *[][]bool) {
		if i < 0 || i >= n || j < 0 || j >= m || (*seen)[i][j] || prevH > heights[i][j] {
			return
		}
		(*seen)[i][j] = true
		if a[i][j] && p[i][j] {
			ans = append(ans, []int{i, j})
		}
		prevH = heights[i][j]
		dfs(i-1, j, prevH, seen)
		dfs(i+1, j, prevH, seen)
		dfs(i, j-1, prevH, seen)
		dfs(i, j+1, prevH, seen)
	}

	for i := 0; i < n; i++ {
		dfs(i, 0, heights[i][0], &p)
		dfs(i, m-1, heights[i][m-1], &a)
	}
	for i := 0; i < m; i++ {
		dfs(0, i, heights[0][i], &p)
		dfs(n-1, i, heights[n-1][i], &a)
	}
	return ans
}

// @lc code=end

