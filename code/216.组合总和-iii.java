import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=216 lang=java
 *
 * [216] 组合总和 III
 *
 * https://leetcode-cn.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (71.70%)
 * Likes:    207
 * Dislikes: 0
 * Total Accepted:    54.6K
 * Total Submissions: 74.3K
 * Testcase Example:  '3\n7'
 *
 * 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
 * 
 * 说明：
 * 
 * 
 * 所有数字都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: k = 3, n = 7
 * 输出: [[1,2,4]]
 * 
 * 
 * 示例 2:
 * 
 * 输入: k = 3, n = 9
 * 输出: [[1,2,6], [1,3,5], [2,3,4]]
 * 
 * 
 */

// @lc code=start
class Solution {
    private int dfs(List<List<Integer>> ans, int k, int n, ArrayList<Integer> combination, int startIdx) {
        if (k == combination.size() && n == 0) {
            ans.add(new ArrayList<>(combination));
            return 1;
        }

        if (combination.size() >= k) {
            return 0;
        }

        if (n <= 0) {
            return -1;
        }

        for(int i = startIdx; i < 10; i++) {
            combination.add(i);
            int res = dfs(ans, k, n - i, combination, i + 1);
            combination.remove(combination.size() - 1);
            if (res == -1) {
                return 0;
            }
        }

        return 0;
    }
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> ans = new ArrayList<>();
        dfs(ans, k, n, new ArrayList<>(), 1);
        return ans;
    }
}
// @lc code=end

