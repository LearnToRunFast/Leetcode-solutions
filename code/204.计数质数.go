/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (38.47%)
 * Likes:    644
 * Dislikes: 0
 * Total Accepted:    135K
 * Total Submissions: 351K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 * 示例 2：
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 * 示例 3：
 *
 * 输入：n = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */

// @lc code=start
// func countPrimes(n int) int {
// 	primes := []int{}
// 	isPrime := make([]bool, n)
// 	for i, _ := range isPrime {
// 		isPrime[i] = true
// 	}
// 	for i := 2; i < n; i++ {
// 		if isPrime[i] {
// 			primes = append(primes, i)
// 		}
// 		for _, v := range primes {
// 			nextIdx := i * v
// 			// i % p == 0 if i is mutiple of prime v
// 			if nextIdx >= n {
// 				break
// 			}
// 			isPrime[nextIdx] = false
// 			if i%v == 0 {
// 				break
// 			}
// 		}
// 	}
// 	return len(primes)
// }
func countPrimes(n int) int {
	isPrimes := make([]bool, n)
	for i := range isPrimes {
		isPrimes[i] = true
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrimes[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrimes[j] = false
			}
		}
	}
	return ans
}

// @lc code=end

