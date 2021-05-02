/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 *
 * https://leetcode-cn.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (55.78%)
 * Likes:    257
 * Dislikes: 0
 * Total Accepted:    19.7K
 * Total Submissions: 35.3K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * 给你一个数组 nums ，请你完成两类查询，其中一类查询要求更新数组下标对应的值，另一类查询要求返回数组中某个范围内元素的总和。
 *
 * 实现 NumArray 类：
 *
 *
 *
 *
 * NumArray(int[] nums) 用整数数组 nums 初始化对象
 * void update(int index, int val) 将 nums[index] 的值更新为 val
 * int sumRange(int left, int right) 返回子数组 nums[left, right] 的总和（即，nums[left] +
 * nums[left + 1], ..., nums[right]）
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["NumArray", "sumRange", "update", "sumRange"]
 * [[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
 * 输出：
 * [null, 9, null, 8]
 *
 * 解释：
 * NumArray numArray = new NumArray([1, 3, 5]);
 * numArray.sumRange(0, 2); // 返回 9 ，sum([1,3,5]) = 9
 * numArray.update(1, 2);   // nums = [1,2,5]
 * numArray.sumRange(0, 2); // 返回 8 ，sum([1,2,5]) = 8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -100
 * 0
 * -100
 * 0
 * 最多调用 3 * 10^4 次 update 和 sumRange 方法
 *
 *
 *
 *
 */

// @lc code=start
type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	m := 2 * n
	tree := make([]int, m)
	for i := range nums {
		tree[i+n] = nums[i]
	}

	// start at 1
	for i := n - 1; i > 0; i-- {
		// parent = left child + right child
		tree[i] = tree[2*i] + tree[2*i+1]
	}

	return NumArray{tree, n}
}

func (this *NumArray) Update(index int, val int) {
	idx := index + this.n
	// update the value first
	this.tree[idx] = val
	for idx > 1 {
		pIdx := idx >> 1
		this.tree[pIdx] = this.tree[2*pIdx] + this.tree[2*pIdx+1]
		idx = pIdx
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	l := left + this.n
	r := right + this.n
	sum := 0
	for l <= r {
		if l%2 == 1 {
			sum += this.tree[l]
			l++
		}
		if r%2 == 0 {
			sum += this.tree[r]
			r--
		}
		l >>= 1
		r >>= 1
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end

