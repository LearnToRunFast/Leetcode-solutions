/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 *
 * https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (56.44%)
 * Likes:    537
 * Dislikes: 0
 * Total Accepted:    81K
 * Total Submissions: 143.4K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * 给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。
 *
 *
 *
 * 示例：
 *
 * 输入：
 * A: [1,2,3,2,1]
 * B: [3,2,1,4,7]
 * 输出：3
 * 解释：
 * 长度最长的公共子数组是 [3, 2, 1] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= len(A), len(B) <= 1000
 * 0 <= A[i], B[i] < 100
 *
 *
 */

// @lc code=start
// rabin karp
const (
	base, mod = 101, 1 << 31
)

// 使用快速幂计算 x^n % mod 的值
func qPow(x, n int) int {
	res := 1
	for n > 0 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		n >>= 1
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func getHash(nums []int, j int) int {
	hash := 0
	for i := 0; i < j; i++ {
		hash = (hash*base + nums[i]) % mod
	}
	return hash
}
func equalArr(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
func check(length int, nums1, nums2 []int) bool {
	hashA := getHash(nums1, length)
	mult := qPow(base, length-1)
	mapA := map[int][]int{hashA: []int{0, length}}
	for i := length; i < len(nums1); i++ {
		toRemove := nums1[i-length] * mult % mod
		// in case if toRemove more than hashA
		hashA = ((hashA-toRemove+mod)%mod*base + nums1[i]) % mod
		mapA[hashA] = []int{i - length + 1, i + 1}
	}

	hashB := getHash(nums2, length)
	if v, ok := mapA[hashB]; ok {
		return equalArr(nums1[v[0]:v[1]], nums2[0:length])
	}
	for i := length; i < len(nums2); i++ {
		toRemove := nums2[i-length] * mult % mod
		hashB = ((hashB-toRemove+mod)%mod*base + nums2[i]) % mod
		if v, ok := mapA[hashB]; ok {
			return equalArr(nums1[v[0]:v[1]], nums2[i-length+1:i+1])
		}
	}
	return false
}
func findLength(nums1 []int, nums2 []int) int {
	// start at the length of half, if there is a match,
	// increase the length by 1/4, if not, decrease the length by 1/4
	left, right := 1, min(len(nums1), len(nums2))+1
	for left < right {
		mid := (left + right) >> 1
		if check(mid, nums1, nums2) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
// func findLength(nums1 []int, nums2 []int) int {
// 	// define dp[i][j] as the longest common subarray of nums1[:i] and nums2[:j]
// 	m, n := len(nums1), len(nums2)
// 	dp := make([][]int, m+1)
// 	for i := range dp {
// 		dp[i] = make([]int, n+1)
// 	}
// 	res := 0
// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			if nums1[i-1] == nums2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 				res = max(res, dp[i][j])
// 			}
// 		}
// 	}
// 	return res
// }

// @lc code=end

