/*
 * @lc app=leetcode.cn id=341 lang=golang
 *
 * [341] 扁平化嵌套列表迭代器
 *
 * https://leetcode-cn.com/problems/flatten-nested-list-iterator/description/
 *
 * algorithms
 * Medium (72.78%)
 * Likes:    332
 * Dislikes: 0
 * Total Accepted:    45.2K
 * Total Submissions: 62.1K
 * Testcase Example:  '[[1,1],2,[1,1]]'
 *
 * 给你一个嵌套的整型列表。请你设计一个迭代器，使其能够遍历这个整型列表中的所有整数。
 *
 * 列表中的每一项或者为一个整数，或者是另一个列表。其中列表的元素也可能是整数或是其他列表。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [[1,1],2,[1,1]]
 * 输出: [1,1,2,1,1]
 * 解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,1,2,1,1]。
 *
 * 示例 2:
 *
 * 输入: [1,[4,[6]]]
 * 输出: [1,4,6]
 * 解释: 通过重复调用 next 直到 hasNext 返回 false，next 返回的元素的顺序应该是: [1,4,6]。
 *
 *
 */

// @lc code=start
type NestedIterator struct {
	stack []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	n := len(nestedList)
	stack := make([]*NestedInteger, n)
	for i := 0; i < n; i++ {
		stack[i] = nestedList[n-1-i]
	}
	return &NestedIterator{stack}
}

func (this *NestedIterator) parse() {

	for len(this.stack) > 0 {
		lastIdx := len(this.stack) - 1
		item := this.stack[lastIdx]
		if item.IsInteger() {
			return
		}
		this.stack = this.stack[:lastIdx]
		list := item.GetList()
		for i := len(list) - 1; i >= 0; i-- {
			this.stack = append(this.stack, list[i])
		}
	}

}
func (this *NestedIterator) Next() int {
	this.parse()
	lastIdx := len(this.stack) - 1
	item := this.stack[lastIdx]
	this.stack = this.stack[:lastIdx]
	return item.GetInteger()

}

func (this *NestedIterator) HasNext() bool {
	this.parse()
	return len(this.stack) > 0
}

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// type NestedIterator struct {
// 	list []int
// }

// func unpack(list *[]int, nestedList []*NestedInteger) {
// 	for _, nextedInteger := range nestedList {
// 		if nextedInteger.IsInteger() {
// 			*list = append(*list, nextedInteger.GetInteger())
// 			continue
// 		}
// 		unpack(list, nextedInteger.GetList())
// 	}
// }

// func Constructor(nestedList []*NestedInteger) *NestedIterator {
// 	list := []int{}
// 	unpack(&list, nestedList)
// 	return &NestedIterator{list}
// }

// func (this *NestedIterator) Next() int {
// 	x := this.list[0]
// 	this.list = this.list[1:]
// 	return x
// }

// func (this *NestedIterator) HasNext() bool {
// 	return len(this.list) > 0
// }
// @lc code=end

