/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 *
 * https://leetcode-cn.com/problems/candy/description/
 *
 * algorithms
 * Hard (48.24%)
 * Likes:    493
 * Dislikes: 0
 * Total Accepted:    66.6K
 * Total Submissions: 138.1K
 * Testcase Example:  '[1,0,2]'
 *
 * 老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
 *
 * 你需要按照以下要求，帮助老师给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 评分更高的孩子必须比他两侧的邻位孩子获得更多的糖果。
 *
 *
 * 那么这样下来，老师至少需要准备多少颗糖果呢？
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[1,0,2]
 * 输出：5
 * 解释：你可以分别给这三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2：
 *
 *
 * 输入：[1,2,2]
 * 输出：4
 * 解释：你可以分别给这三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
 *
 */

// @lc code=start
package main

// func max(a, b int) int {
// 	if a >= b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
//
// func candy(ratings []int) int {
// 	n := len(ratings)
// 	if n < 2 {
// 		return n
// 	}
// 	left := make([]int, n, n)
// 	left[0] = 1
// 	for i := 1; i < n; i++ {
// 		if ratings[i] > ratings[i-1] {
// 			left[i] = left[i-1] + 1
// 		} else {
// 			left[i] = 1
// 		}
// 	}
// 	right := 1
// 	sum := left[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		if ratings[i] > ratings[i+1] {
// 			right++
// 		} else {
// 			right = 1
// 		}
// 		sum += max(right, left[i])
// 	}

// 	return sum
// }
func candy(ratings []int) int {
	n := len(ratings)
	sum, pre, incLen, decLen := 1, 1, 1, 0
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			decLen = 0
			pre++
			sum += pre
			incLen = pre
		} else if ratings[i] == ratings[i-1] {
			decLen = 0
			pre = 1
			sum++
			incLen = 1
		} else {
			decLen++
			pre = 1
			// if decLen is >= incLen, we need to include the top most value
			if decLen == incLen {
				decLen++
			}
			sum += decLen
		}
	}
	return sum
}

// @lc code=end
