/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (63.39%)
 * Likes:    597
 * Dislikes: 0
 * Total Accepted:    69.8K
 * Total Submissions: 109.9K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * 给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
 * 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
 * 输出：13
 * 解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix = [[-5]], k = 1
 * 输出：-5
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == matrix.length
 * n == matrix[i].length
 * 1
 * -10^9
 * 题目数据 保证 matrix 中的所有行和列都按 非递减顺序 排列
 * 1
 *
 *
 */

// @lc code=start
// type keyValue struct {
// 	row   int
// 	col   int
// 	value int
// }
// type MinHeap []*keyValue

// func (m MinHeap) Len() int           { return len(m) }
// func (m MinHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
// func (m MinHeap) Less(i, j int) bool { return m[i].value < m[j].value }
// func (m *MinHeap) Push(x interface{}) {
// 	*m = append(*m, x.(*keyValue))
// }
// func (m *MinHeap) Pop() interface{} {
// 	old := *m
// 	n := len(old)
// 	x := old[n-1]
// 	*m = old[:n-1]
// 	return x
// }
// func max(i, j int) int {
// 	if i < j {
// 		return j
// 	}
// 	return i
// }
// func kthSmallest(matrix [][]int, k int) int {
// 	n := len(matrix)
// 	if n == 0 {
// 		return 0
// 	}
// 	m := len(matrix[0])
// 	sum := m * n
// 	if k > sum {
// 		return matrix[n-1][m-1]
// 	}

// 	h := &MinHeap{}
// 	for i, row := range matrix {
// 		heap.Push(h, &keyValue{i, 0, row[0]})
// 	}
// 	ans := 0
// 	for k > 0 {
// 		x := heap.Pop(h).(*keyValue)

// 		newCol := x.col + 1
// 		if newCol < m {
// 			heap.Push(h, &keyValue{x.row, newCol, matrix[x.row][newCol]})
// 		}
// 		ans = x.value
// 		k--
// 	}
// 	return ans
// }
// method 2
func isLessK(matrix [][]int, n, k, mid int) bool {
	i, j := n-1, 0
	nums := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			nums += i + 1
			j++
		} else {
			i--
		}
	}
	return nums < k
}
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	sum := m * n
	if k > sum {
		return matrix[n-1][m-1]
	}
	lo, hi := matrix[0][0], matrix[n-1][m-1]
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isLessK(matrix, n, k, mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo

}

// @lc code=end

