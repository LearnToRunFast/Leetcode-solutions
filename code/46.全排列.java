import java.awt.List;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;

/*
 * @lc app=leetcode.cn id=46 lang=java
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (76.70%)
 * Likes:    910
 * Dislikes: 0
 * Total Accepted:    198.3K
 * Total Submissions: 257.8K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */

// @lc code=start

// // first verion check duplicate
// class Solution {
//     private void dfs(List<List<Integer>> ans, int[] nums, List<Integer> combination) {
//         if (combination.size() == nums.length) {
//             ans.add(new ArrayList<>(combination));
//             return;
//         }

//         for(int i = 0; i < nums.length; i++) {
//             if (combination.contains(nums[i])) continue;
//             combination.add(nums[i]);
//             dfs(ans, nums, combination);
//             combination.remove(combination.size() - 1);
//         }

//     }
//     public List<List<Integer>> permute(int[] nums) {
//         List<List<Integer>> ans = new ArrayList<>();

//         if (nums == null || nums.length == 0) return ans;
//         dfs(ans, nums, new ArrayList<>());

//         return ans;
//     }
// }


// second verion swap
class Solution {
    private void dfs(List<List<Integer>> ans, int n, List<Integer> combination, int steps) {
        if (steps == n) {
            ans.add(new ArrayList<Integer>(combination));
            return;
        }

        for(int i = steps; i < n; i++) {
            Collections.swap(combination, steps, i);
            dfs(ans, n, combination, steps + 1);
            Collections.swap(combination, steps, i);
        }

    }
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> ans = new ArrayList<>();
        List<Integer> combination = new ArrayList<>();
        if (nums == null || nums.length == 0) return ans;

        for(int num : nums) {
            combination.add(num);
        }

        dfs(ans, nums.length, combination, 0);

        return ans;
    }
}
// @lc code=end

