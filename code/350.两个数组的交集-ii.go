/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (54.27%)
 * Likes:    489
 * Dislikes: 0
 * Total Accepted:    196.2K
 * Total Submissions: 361.2K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2,2]
 *
 *
 * 示例 2:
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[4,9]
 *
 *
 *
 * 说明：
 *
 *
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 * 进阶：
 *
 *
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 *
 *
 */

// @lc code=start
func getMin(i, j int) int {
	if i > j {
		return j
	}
	return i
}
func intersect(nums1 []int, nums2 []int) []int {
	dist1, dist2 := map[int]int{}, map[int]int{}
	for _, num := range nums1 {
		dist1[num]++
	}
	for _, num := range nums2 {
		dist2[num]++
	}
	ans := []int{}
	for k, v := range dist1 {
		if u, ok := dist2[k]; ok {
			min := getMin(v, u)
			for min > 0 {
				ans = append(ans, k)
				min--
			}
		}
	}
	return ans
}

// @lc code=end

