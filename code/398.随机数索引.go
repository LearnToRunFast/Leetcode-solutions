/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 *
 * https://leetcode-cn.com/problems/random-pick-index/description/
 *
 * algorithms
 * Medium (63.97%)
 * Likes:    109
 * Dislikes: 0
 * Total Accepted:    12.2K
 * Total Submissions: 18.9K
 * Testcase Example:  '["Solution","pick","pick","pick"]\n[[[1,2,3,3,3]],[3],[1],[3]]'
 *
 * 给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。
 *
 * 注意：
 * 数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。
 *
 * 示例:
 *
 *
 * int[] nums = new int[] {1,2,3,3,3};
 * Solution solution = new Solution(nums);
 *
 * // pick(3) 应该返回索引 2,3 或者 4。每个索引的返回概率应该相等。
 * solution.pick(3);
 *
 * // pick(1) 应该返回 0。因为只有nums[0]等于1。
 * solution.pick(1);
 *
 *
 */

// @lc code=start
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}
func (s *Solution) Pick(target int) int {
	n := len(s.nums)
	ans := 0
	count := 0
	for i := 0; i < n; i++ {
		if s.nums[i] == target {
			if rand.Intn(count+1) == count {
				ans = i
			}
			count++
		}

	}
	return ans

}

// type Solution struct {
// 	dict map[int][]int
// }

// func Constructor(nums []int) Solution {
// 	dict := map[int][]int{} // key is value, value is array of indexs
// 	for i, num := range nums {
// 		if _, ok := dict[num]; !ok {
// 			dict[num] = []int{}
// 		}
// 		dict[num] = append(dict[num], i)
// 	}
// 	return Solution{dict}
// }

// func (this *Solution) Pick(target int) int {
// 	arr := this.dict[target]
// 	n := len(arr)
// 	if n == 0 {
// 		return -1
// 	}
// 	num := rand.Intn(n)
// 	return arr[num]
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

