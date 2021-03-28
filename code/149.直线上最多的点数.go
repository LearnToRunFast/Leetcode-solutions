/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 *
 * https://leetcode-cn.com/problems/max-points-on-a-line/description/
 *
 * algorithms
 * Hard (23.73%)
 * Likes:    229
 * Dislikes: 0
 * Total Accepted:    20.6K
 * Total Submissions: 86K
 * Testcase Example:  '[[1,1],[2,2],[3,3]]'
 *
 * 给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
 *
 * 示例 1:
 *
 * 输入: [[1,1],[2,2],[3,3]]
 * 输出: 3
 * 解释:
 * ^
 * |
 * |        o
 * |     o
 * |  o
 * +------------->
 * 0  1  2  3  4
 *
 *
 * 示例 2:
 *
 * 输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
 * 输出: 4
 * 解释:
 * ^
 * |
 * |  o
 * |     o        o
 * |        o
 * |  o        o
 * +------------------->
 * 0  1  2  3  4  5  6
 *
 */

// @lc code=start
package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	ans := 0
	for i, v := range points {
		if ans >= n-i {
			break
		}
		lookup := make(map[string]int)
		dup := 0
		currMax := 0
		for _, u := range points[i+1:] {
			x := v[0] - u[0]
			y := v[1] - u[1]
			if x == 0 && y == 0 {
				dup++
				continue
			}
			gcf := gcd(x, y)
			x = x / gcf
			y = y / gcf
			key := fmt.Sprintf("%s:%s", x, y)
			lookup[key]++
			currMax = max(currMax, lookup[key])
		}
		ans = max(ans, currMax+dup+1)
	}
	return ans
}

// @lc code=end
