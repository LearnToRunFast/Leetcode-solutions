/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 *
 * https://leetcode-cn.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (56.22%)
 * Likes:    79
 * Dislikes: 0
 * Total Accepted:    19.8K
 * Total Submissions: 35.2K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 * 给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold
 * Medal", "Silver Medal", "Bronze Medal"）。
 *
 * (注：分数越高的选手，排名越靠前。)
 *
 * 示例 1:
 *
 *
 * 输入: [5, 4, 3, 2, 1]
 * 输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * 解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal"
 * and "Bronze Medal").
 * 余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
 *
 * 提示:
 *
 *
 * N 是一个正整数并且不会超过 10000。
 * 所有运动员的成绩都不相同。
 *
 *
 */

// @lc code=start
type ScoreIdx struct {
	score int
	idx   int
}

func findRelativeRanks(score []int) []string {
	n := len(score)
	ans := make([]string, n)
	scoreIdxs := make([]ScoreIdx, n)
	for i, s := range score {
		scoreIdxs[i] = ScoreIdx{s, i}
	}
	sort.Slice(scoreIdxs, func(i, j int) bool {
		return scoreIdxs[i].score > scoreIdxs[j].score
	})

	Prize := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	for i, scoreIdx := range scoreIdxs {
		if i < 3 {
			ans[scoreIdx.idx] = Prize[i]
			continue
		}
		ans[scoreIdx.idx] = strconv.Itoa(i + 1)
	}
	return ans
}

// @lc code=end

