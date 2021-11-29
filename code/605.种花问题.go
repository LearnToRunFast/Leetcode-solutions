/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 *
 * https://leetcode-cn.com/problems/can-place-flowers/description/
 *
 * algorithms
 * Easy (33.32%)
 * Likes:    407
 * Dislikes: 0
 * Total Accepted:    114.3K
 * Total Submissions: 343.4K
 * Testcase Example:  '[1,0,0,0,1]\n1'
 *
 * 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
 *
 * 给你一个整数数组  flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n
 * ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：flowerbed = [1,0,0,0,1], n = 1
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：flowerbed = [1,0,0,0,1], n = 2
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * flowerbed[i] 为 0 或 1
 * flowerbed 中不存在相邻的两朵花
 * 0
 *
 *
 */

// @lc code=start
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	for i := 0; i < m; i++ {
		v := flowerbed[i]
		if v == 0 && (i+1 == m || flowerbed[i+1] == 0) {
			n--
			i++
		} else if v == 1 {
			i++
		}
	}
	return n <= 0
}

// @lc code=end

