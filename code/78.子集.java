/*
 * @lc app=leetcode.cn id=78 lang=java
 *
 * [78] 子集
 *
 * https://leetcode-cn.com/problems/subsets/description/
 *
 * algorithms
 * Medium (79.22%)
 * Likes:    879
 * Dislikes: 0
 * Total Accepted:    170.4K
 * Total Submissions: 215.1K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
 * 
 * 说明：解集不能包含重复的子集。
 * 
 * 示例:
 * 
 * 输入: nums = [1,2,3]
 * 输出:
 * [
 * ⁠ [3],
 * [1],
 * [2],
 * [1,2,3],
 * [1,3],
 * [2,3],
 * [1,2],
 * []
 * ]
 * 
 */

// @lc code=start
// backtrack
// class Solution {
//     private void backtrack(int step, int[] nums, List<Integer> tmp, List<List<Integer>> ans) {
        
//         if (step == nums.length) {
//             ans.add(new ArrayList<>(tmp));
//             return;
//         }
//         tmp.add(nums[step]);
//         backtrack(step + 1, nums, tmp, ans);
//         tmp.remove(tmp.size() - 1);
//         backtrack(step + 1, nums, tmp, ans);
//     }
//     public List<List<Integer>> subsets(int[] nums) {
//         List<List<Integer>> ans = new ArrayList<>();
//         List<Integer> tmp = new ArrayList<>();
//         backtrack(0, nums, tmp, ans);
//         return ans;
//     }
// }
class Solution {

    public List<List<Integer>> subsets(int[] nums) {
        
        List<List<Integer>> ans = new ArrayList<>();
        List<Integer> tmp = new ArrayList<>();

        int n = nums.length;
        int m = 1 << n;
        for (int mask = 0; mask < m; mask++) {

            for (int i = 0; i < n; i++) {
                if ((mask & (1 << i)) != 0) {
                    tmp.add(nums[i]);
                }
            }
            ans.add(new ArrayList<>(tmp));
            tmp = new ArrayList<>();
        }
        return ans;
    }
}
// @lc code=end

