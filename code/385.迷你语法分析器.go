/*
 * @lc app=leetcode.cn id=385 lang=golang
 *
 * [385] 迷你语法分析器
 *
 * https://leetcode-cn.com/problems/mini-parser/description/
 *
 * algorithms
 * Medium (41.18%)
 * Likes:    63
 * Dislikes: 0
 * Total Accepted:    6.5K
 * Total Submissions: 15.8K
 * Testcase Example:  '"324"'
 *
 * 给定一个用字符串表示的整数的嵌套列表，实现一个解析它的语法分析器。
 *
 * 列表中的每个元素只可能是整数或整数嵌套列表
 *
 * 提示：你可以假定这些字符串都是格式良好的：
 *
 *
 * 字符串非空
 * 字符串不包含空格
 * 字符串只包含数字0-9、[、-、,、]
 *
 *
 *
 *
 * 示例 1：
 *
 * 给定 s = "324",
 *
 * 你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
 *
 *
 * 示例 2：
 *
 * 给定 s = "[123,[456,[789]]]",
 *
 * 返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
 *
 * 1. 一个 integer 包含值 123
 * 2. 一个包含两个元素的嵌套列表：
 * ⁠   i.  一个 integer 包含值 456
 * ⁠   ii. 一个包含一个元素的嵌套列表
 * ⁠        a. 一个 integer 包含值 789
 *
 *
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func makeIntObj(num []byte) *NestedInteger {
	temp := NestedInteger{}
	val, _ := strconv.Atoi(string(num))
	temp.SetInteger(val)
	return &temp
}
func deserialize(s string) *NestedInteger {
	stack := make([]NestedInteger, 0)
	num := []byte{}
	for i := range s {
		switch s[i] {
		case '[':
			stack = append(stack, NestedInteger{})
		case ']':
			if len(num) > 0 {
				lastIdx := len(stack) - 1
				stack[lastIdx].Add(*makeIntObj(num))
				num = []byte{}
			}
			n := len(stack)
			if n > 1 {
				last := stack[n-1]
				stack = stack[:n-1]
				stack[n-2].Add(last)
			}
		case ',':
			if len(num) > 0 {
				lastIdx := len(stack) - 1
				stack[lastIdx].Add(*makeIntObj(num))
				num = []byte{}
			}
		default:
			num = append(num, s[i])
		}
	}
	if len(num) > 0 {
		return makeIntObj(num)
	}
	return &stack[0]
}

// @lc code=end

