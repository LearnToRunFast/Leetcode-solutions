/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 *
 * https://leetcode-cn.com/problems/wiggle-sort-ii/description/
 *
 * algorithms
 * Medium (37.31%)
 * Likes:    251
 * Dislikes: 0
 * Total Accepted:    20.9K
 * Total Submissions: 55.9K
 * Testcase Example:  '[1,5,1,1,6,4]'
 *
 * 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
 *
 * 你可以假设所有输入数组都可以得到满足题目要求的结果。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,5,1,1,6,4]
 * 输出：[1,6,1,5,1,4]
 * 解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3,2,2,3,1]
 * 输出：[2,3,1,3,1,2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 0
 * 题目数据保证，对于给定的输入 nums ，总能产生满足题目要求的结果
 *
 *
 *
 *
 * 进阶：你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
 *
 */

// @lc code=start
// func shuffle(nums *[]int) {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := range *nums {
// 		r := rand.Intn(i + 1)
// 		(*nums)[i], (*nums)[r] = (*nums)[r], (*nums)[i]
// 	}
// }
// func partition(nums *[]int, lo, hi int) (int, int) {
// 	lt, i, gt := lo, lo+1, hi
// 	v := (*nums)[lo]
// 	for i <= gt {
// 		if (*nums)[i] > v {
// 			(*nums)[gt], (*nums)[i] = (*nums)[i], (*nums)[gt]
// 			gt--
// 		} else if (*nums)[i] < v {
// 			(*nums)[lt], (*nums)[i] = (*nums)[i], (*nums)[lt]
// 			lt++
// 			i++
// 		} else {
// 			i++
// 		}
// 	}
// 	return lt, gt
// }
// func quickSelect(nums *[]int, k int) {
// 	n := len(*nums)
// 	if k < 0 || k >= n {
// 		return
// 	}
// 	shuffle(nums)
// 	lo, hi := 0, n-1
// 	for lo < hi {
// 		lt, gt := partition(nums, lo, hi)
// 		if k > gt {
// 			lo = gt + 1
// 		} else if k < lt {
// 			hi = lt - 1
// 		} else {
// 			return
// 		}
// 	}
// 	return
// }

// func wiggleSort(nums []int) {
// 	n := len(nums)
// 	mid := n / 2
// 	quickSelect(&nums, mid)
// 	if n%2 == 0 {
// 		mid--
// 	}
// 	temp := []int{}
// 	temp = append(temp, nums...)
// 	l, r := mid, n-1
// 	i := 0
// 	for i < n {
// 		if l >= 0 {
// 			nums[i] = temp[l]
// 			l--
// 			i++
// 		}
// 		if r > mid {
// 			nums[i] = temp[r]
// 			i++
// 			r--
// 		}
// 	}
// }
func wiggleSort(nums []int) {
	n := len(nums)
	bucket := make([]int, 5001)
	for _, v := range nums {
		bucket[v]++
	}
	curr := 5000
	for i := 1; i < n; {
		if bucket[curr] == 0 {
			curr--
			continue
		}
		nums[i] = curr
		bucket[curr]--
		i += 2
	}
	for i := 0; i < n && curr >= 0; {
		if bucket[curr] == 0 {
			curr--
			continue
		}
		nums[i] = curr
		bucket[curr]--
		i += 2
	}
}

// @lc code=end

