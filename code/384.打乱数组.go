/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 *
 * https://leetcode-cn.com/problems/shuffle-an-array/description/
 *
 * algorithms
 * Medium (56.19%)
 * Likes:    136
 * Dislikes: 0
 * Total Accepted:    40.9K
 * Total Submissions: 72.4K
 * Testcase Example:  '["Solution","shuffle","reset","shuffle"]\n[[[1,2,3]],[],[],[]]'
 *
 * 给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。
 *
 * 实现 Solution class:
 *
 *
 * Solution(int[] nums) 使用整数数组 nums 初始化对象
 * int[] reset() 重设数组到它的初始状态并返回
 * int[] shuffle() 返回数组随机打乱后的结果
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入
 * ["Solution", "shuffle", "reset", "shuffle"]
 * [[[1, 2, 3]], [], [], []]
 * 输出
 * [null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
 *
 * 解释
 * Solution solution = new Solution([1, 2, 3]);
 * solution.shuffle();    // 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。例如，返回
 * [3, 1, 2]
 * solution.reset();      // 重设数组到它的初始状态 [1, 2, 3] 。返回 [1, 2, 3]
 * solution.shuffle();    // 随机返回数组 [1, 2, 3] 打乱后的结果。例如，返回 [1, 3, 2]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10^6
 * nums 中的所有元素都是 唯一的
 * 最多可以调用 5 * 10^4 次 reset 和 shuffle
 *
 *
 */

// @lc code=start
type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.original
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	temp := append([]int{}, s.original...)
	for i := range temp {
		r := rand.Intn(i + 1)
		temp[i], temp[r] = temp[r], temp[i]
	}
	return temp
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

