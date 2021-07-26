/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最少移动次数使数组元素相等 II
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii/description/
 *
 * algorithms
 * Medium (59.97%)
 * Likes:    133
 * Dislikes: 0
 * Total Accepted:    12.8K
 * Total Submissions: 21.4K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。 您可以假设数组的长度最多为10000。
 *
 * 例如:
 *
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 2
 *
 * 说明：
 * 只有两个动作是必要的（记得每一步仅可使其中一个元素加1或减1）：
 *
 * [1,2,3]  =>  [2,2,3]  =>  [2,2,2]
 *
 *
 */

// @lc code=start
// func minMoves2(nums []int) int {
// 	sort.Ints(nums)
// 	var sum = 0
// 	mid := nums[len(nums)/2]
// 	for _, num := range nums {
// 		sum += int(math.Abs(float64(num - mid)))
// 	}
// 	return sum
// }
func minMoves2(nums []int) int {
	idx := len(nums) / 2
	quickSelect(&nums, idx)
	mid := nums[idx]
	var sum = 0
	for _, num := range nums {
		sum += int(math.Abs(float64(num - mid)))
	}
	return sum
}
func shuffle(nums *[]int) {
	rand.Seed(time.Now().UnixNano())
	for i := range *nums {
		r := rand.Intn(i + 1)
		(*nums)[i], (*nums)[r] = (*nums)[r], (*nums)[i]
	}
}
func partition(nums *[]int, lo, hi int) (int, int) {
	lt, i, gt := lo, lo+1, hi
	v := (*nums)[lo]
	for i <= gt {
		if (*nums)[i] > v {
			(*nums)[gt], (*nums)[i] = (*nums)[i], (*nums)[gt]
			gt--
		} else if (*nums)[i] < v {
			(*nums)[lt], (*nums)[i] = (*nums)[i], (*nums)[lt]
			lt++
			i++
		} else {
			i++
		}
	}
	return lt, gt
}
func quickSelect(nums *[]int, k int) {
	n := len(*nums)
	if k < 0 || k >= n {
		return
	}
	shuffle(nums)
	lo, hi := 0, n-1
	for lo < hi {
		lt, gt := partition(nums, lo, hi)
		if k > gt {
			lo = gt + 1
		} else if k < lt {
			hi = lt - 1
		} else {
			return
		}
	}
	return
}

// @lc code=end

