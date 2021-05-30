/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 *
 * https://leetcode-cn.com/problems/lexicographical-numbers/description/
 *
 * algorithms
 * Medium (73.61%)
 * Likes:    177
 * Dislikes: 0
 * Total Accepted:    18.3K
 * Total Submissions: 24.9K
 * Testcase Example:  '13'
 *
 * 给定一个整数 n, 返回从 1 到 n 的字典顺序。
 *
 * 例如，
 *
 * 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。
 *
 * 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。
 *
 */

// @lc code=start

// func lexicalOrder(n int) []int {
// 	ans := make([]int, n)
// 	count := 0
// 	var addNum func(num int)

// 	addNum = func(num int) {
// 		if num > n {
// 			return
// 		}
// 		ans[count] = num
// 		count++
// 		for i := 0; i <= 9; i++ {
// 			addNum(num*10 + i)
// 		}
// 	}

// 	for i := 1; i <= 9; i++ {
// 		addNum(i)
// 	}
// 	return ans

// }

func lexicalOrder(n int) []int {
	ans := make([]int, n)
	count := 0
	stack := []int{}
	for i := 9; i > 0; i-- {
		if n >= i {
			stack = append(stack, i)
		}
	}
	lastIdx := len(stack) - 1
	for lastIdx >= 0 {
		num := stack[lastIdx]
		ans[count] = num
		count++
		stack = stack[:lastIdx]
		for i := 9; i >= 0; i-- {
			newNum := num*10 + i
			if newNum <= n {
				stack = append(stack, newNum)
			}
		}
		lastIdx = len(stack) - 1
	}

	return ans

}

// @lc code=end

