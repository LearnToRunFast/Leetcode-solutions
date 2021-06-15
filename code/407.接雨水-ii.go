/*
 * @lc app=leetcode.cn id=407 lang=golang
 *
 * [407] 接雨水 II
 *
 * https://leetcode-cn.com/problems/trapping-rain-water-ii/description/
 *
 * algorithms
 * Hard (47.55%)
 * Likes:    346
 * Dislikes: 0
 * Total Accepted:    8.6K
 * Total Submissions: 18K
 * Testcase Example:  '[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]'
 *
 * 给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
 *
 *
 *
 * 示例：
 *
 * 给出如下 3x6 的高度图:
 * [
 * ⁠ [1,4,3,1,3,2],
 * ⁠ [3,2,1,3,2,4],
 * ⁠ [2,3,3,2,3,1]
 * ]
 *
 * 返回 4 。
 *
 *
 *
 *
 * 如上图所示，这是下雨前的高度图[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] 的状态。
 *
 *
 *
 *
 *
 * 下雨后，雨水将会被存储在这些方块中。总的接雨水量是4。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= m, n <= 110
 * 0 <= heightMap[i][j] <= 20000
 *
 *
 */

// @lc code=start
type Cell struct {
	r int
	c int
	h int
}
type MinHeap []Cell

func (m MinHeap) Len() int           { return len(m) }
func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Less(i, j int) bool { return m[i].h < m[j].h }
func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(Cell))
}
func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func trapRainWater(heightMap [][]int) int {
	n := len(heightMap)
	if n == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	m := len(heightMap[0])
	h := MinHeap{}
	heap.Init(&h)

	// add surrounding to heap
	for i, v := range heightMap {
		for j, u := range v {
			if i == 0 || i == n-1 || j == 0 || j == m-1 {
				heap.Push(&h, Cell{i, j, u})
				heightMap[i][j] = -1
			}
		}
	}

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0
	for h.Len() > 0 {
		cell := heap.Pop(&h).(Cell)
		for _, d := range dir {
			r := cell.r + d[0]
			c := cell.c + d[1]
			if r >= 0 && r < n && c >= 0 && c < m && heightMap[r][c] != -1 {
				height := heightMap[r][c]
				diff := cell.h - height
				if diff > 0 {
					ans += diff
					height = cell.h
				}
				heap.Push(&h, Cell{r, c, height})
				heightMap[r][c] = -1 // visited
			}
		}
	}
	return ans
}

// @lc code=end

