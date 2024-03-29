/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 *
 * https://leetcode-cn.com/problems/reconstruct-itinerary/description/
 *
 * algorithms
 * Medium (44.55%)
 * Likes:    402
 * Dislikes: 0
 * Total Accepted:    32.4K
 * Total Submissions: 72.7K
 * Testcase Example:  '[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]'
 *
 * 给定一个机票的字符串二维数组 [from,
 * to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从
 * JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。
 *
 *
 *
 * 提示：
 *
 *
 * 如果存在多种有效的行程，请你按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"]
 * 相比就更小，排序更靠前
 * 所有的机场都用三个大写字母表示（机场代码）。
 * 假定所有机票至少存在一种合理的行程。
 * 所有的机票必须都用一次 且 只能用一次。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
 * 输出：["JFK", "MUC", "LHR", "SFO", "SJC"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
 * 输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
 * 解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"]。但是它自然排序更大更靠后。
 *
 */

// @lc code=start
func dfs(ans *[]string, edges map[string][]string, start string) {
	for {
		if v, ok := edges[start]; !ok || len(v) == 0 {
			break
		}
		next := edges[start][0]
		edges[start] = edges[start][1:]
		dfs(ans, edges, next)
	}
	*ans = append(*ans, start)
}

// according to the question, it says there must be a valid path
func findItinerary(tickets [][]string) []string {

	ans := make([]string, 0)
	edges := map[string][]string{}
	for _, ticket := range tickets {
		edges[ticket[0]] = append(edges[ticket[0]], ticket[1])
	}
	for key := range edges {
		sort.Strings(edges[key])
	}
	dfs(&ans, edges, "JFK")

	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans

}

// @lc code=end

