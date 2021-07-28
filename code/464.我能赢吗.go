/*
 * @lc app=leetcode.cn id=464 lang=golang
 *
 * [464] 我能赢吗
 *
 * https://leetcode-cn.com/problems/can-i-win/description/
 *
 * algorithms
 * Medium (35.33%)
 * Likes:    239
 * Dislikes: 0
 * Total Accepted:    10K
 * Total Submissions: 28.1K
 * Testcase Example:  '10\n11'
 *
 * 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，先使得累计整数和达到或超过 100 的玩家，即为胜者。
 *
 * 如果我们将游戏规则改为 “玩家不能重复使用整数” 呢？
 *
 * 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
 *
 * 给定一个整数 maxChoosableInteger （整数池中可选择的最大数）和另一个整数
 * desiredTotal（累计和），判断先出手的玩家是否能稳赢（假设两位玩家游戏时都表现最佳）？
 *
 * 你可以假设 maxChoosableInteger 不会大于 20， desiredTotal 不会大于 300。
 *
 * 示例：
 *
 * 输入：
 * maxChoosableInteger = 10
 * desiredTotal = 11
 *
 * 输出：
 * false
 *
 * 解释：
 * 无论第一个玩家选择哪个整数，他都会失败。
 * 第一个玩家可以选择从 1 到 10 的整数。
 * 如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
 * 第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
 * 同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。
 *
 *
 */

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	if maxChoosableInteger >= desiredTotal {
		return true
	}

	dp := make([]byte, 1<<(maxChoosableInteger+1))

	var backtracking func(v, left int) bool
	backtracking = func(v, left int) bool {
		if dp[v] != 0 {
			return dp[v] == 1
		}
		dp[v] = 2
		for i := 1; i <= maxChoosableInteger; i++ {
			currBit := 1 << i
			if currBit&v != 0 {
				continue
			}
			if i >= left {
				dp[v] = 1
				break
			}
			secondPlayerWin := backtracking(v|currBit, left-i)
			if !secondPlayerWin {
				dp[v] = 1
				break
			}
		}
		return dp[v] == 1
	}
	return backtracking(0, desiredTotal)
}

// @lc code=end

