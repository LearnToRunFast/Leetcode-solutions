/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 *
 * https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/description/
 *
 * algorithms
 * Hard (41.96%)
 * Likes:    575
 * Dislikes: 0
 * Total Accepted:    44.3K
 * Total Submissions: 105.6K
 * Testcase Example:  '[5,2,6,1]'
 *
 * 给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于
 * nums[i] 的元素的数量。
 *
 *
 *
 * 示例：
 *
 * 输入：nums = [5,2,6,1]
 * 输出：[2,1,1,0]
 * 解释：
 * 5 的右侧有 2 个更小的元素 (2 和 1)
 * 2 的右侧仅有 1 个更小的元素 (1)
 * 6 的右侧有 1 个更小的元素 (1)
 * 1 的右侧有 0 个更小的元素
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start
func countSmaller(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	if n == 0 {
		return ans
	}
	indexs := make([]int, n)
	temp := make([]int, n)

	for i := range indexs {
		indexs[i] = i
	}

	var sort func(l, r int)
	var merge func(l, mid, r int)
	merge = func(l, mid, r int) {
		for i := l; i <= r; i++ {
			temp[i] = indexs[i]
		}
		i := l
		j := mid + 1
		for k := l; k <= r; k++ {
			if i > mid {
				indexs[k] = temp[j]
				j++
			} else if j > r {
				indexs[k] = temp[i]
				i++
				ans[indexs[k]] += r - mid
			} else if nums[temp[i]] <= nums[temp[j]] {
				indexs[k] = temp[i]
				i++
				ans[indexs[k]] += j - mid - 1
			} else {
				indexs[k] = temp[j]
				j++
			}
		}
	}
	sort = func(l, r int) {
		if l >= r {
			return
		}
		mid := l + (r-l)/2
		sort(l, mid)
		sort(mid+1, r)
		if nums[indexs[mid]] <= nums[indexs[mid+1]] {
			return
		}
		merge(l, mid, r)
	}
	sort(0, n-1)
	return ans
}

// @lc code=end

