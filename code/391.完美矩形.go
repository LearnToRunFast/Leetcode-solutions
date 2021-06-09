/*
 * @lc app=leetcode.cn id=391 lang=golang
 *
 * [391] 完美矩形
 *
 * https://leetcode-cn.com/problems/perfect-rectangle/description/
 *
 * algorithms
 * Hard (34.07%)
 * Likes:    81
 * Dislikes: 0
 * Total Accepted:    3.8K
 * Total Submissions: 11.1K
 * Testcase Example:  '[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]'
 *
 * 我们有 N 个与坐标轴对齐的矩形, 其中 N > 0, 判断它们是否能精确地覆盖一个矩形区域。
 *
 * 每个矩形用左下角的点和右上角的点的坐标来表示。例如， 一个单位正方形可以表示为 [1,1,2,2]。 ( 左下角的点的坐标为 (1, 1)
 * 以及右上角的点的坐标为 (2, 2) )。
 *
 *
 *
 * 示例 1:
 *
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4],
 * ⁠ [1,3,2,4],
 * ⁠ [2,3,3,4]
 * ]
 *
 * 返回 true。5个矩形一起可以精确地覆盖一个矩形区域。
 *
 *
 *
 *
 *
 *
 * 示例 2:
 *
 * rectangles = [
 * ⁠ [1,1,2,3],
 * ⁠ [1,3,2,4],
 * ⁠ [3,1,4,2],
 * ⁠ [3,2,4,4]
 * ]
 *
 * 返回 false。两个矩形之间有间隔，无法覆盖成一个矩形。
 *
 *
 *
 *
 *
 *
 * 示例 3:
 *
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [3,2,4,4]
 * ]
 *
 * 返回 false。图形顶端留有间隔，无法覆盖成一个矩形。
 *
 *
 *
 *
 *
 *
 * 示例 4:
 *
 * rectangles = [
 * ⁠ [1,1,3,3],
 * ⁠ [3,1,4,2],
 * ⁠ [1,3,2,4],
 * ⁠ [2,2,4,4]
 * ]
 *
 * 返回 false。因为中间有相交区域，虽然形成了矩形，但不是精确覆盖。
 *
 *
 */

// @lc code=start

func isRectangleCover(rectangles [][]int) bool {
	X1, Y1, X2, Y2 := 1<<31-1, 1<<31-1, -1<<31, -1<<31

	area := 0
	dict := map[string]int{}
	for _, r := range rectangles {
		x1, y1, x2, y2 := r[0], r[1], r[2], r[3]
		if X1 > x1 {
			X1 = x1
		}
		if Y1 > y1 {
			Y1 = y1
		}
		if X2 < x2 {
			X2 = x2
		}
		if Y2 < y2 {
			Y2 = y2
		}
		area += (y2 - y1) * (x2 - x1)
		dict[fmt.Sprintf("%v-%v", x1, y1)]++
		dict[fmt.Sprintf("%v-%v", x2, y1)]++
		dict[fmt.Sprintf("%v-%v", x1, y2)]++
		dict[fmt.Sprintf("%v-%v", x2, y2)]++
	}
	if area != (Y2-Y1)*(X2-X1) {
		return false
	}
	dict[fmt.Sprintf("%v-%v", X1, Y1)]++
	dict[fmt.Sprintf("%v-%v", X2, Y1)]++
	dict[fmt.Sprintf("%v-%v", X1, Y2)]++
	dict[fmt.Sprintf("%v-%v", X2, Y2)]++
	for _, v := range dict {
		if v%2 == 1 {
			return false
		}
	}
	return true
}

// @lc code=end

