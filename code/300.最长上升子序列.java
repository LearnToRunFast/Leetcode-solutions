import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=300 lang=java
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (45.16%)
 * Likes:    1109
 * Dislikes: 0
 * Total Accepted:    161K
 * Total Submissions: 355.6K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 * 
 * 示例:
 * 
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4 
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 * 
 * 说明:
 * 
 * 
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 * 
 * 
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 * 
 */

// @lc code=start
// dp
// class Solution {
//     public int lengthOfLIS(int[] nums) {
//         if (nums == null || nums.length == 0) return 0;
//         int dp[] = new int[nums.length];
//         Arrays.fill(dp, 1);
//         int ans = 0;
//         for (int i = 0; i < nums.length; i++) {
//             for (int j = 0; j < i; j++) {
//                 if (nums[j] < nums[i]) {
//                     dp[i] = Math.max(dp[i], dp[j] + 1);
//                 }
//             }
//             ans = Math.max(ans, dp[i]);
//         }
//         return ans;
//     }
// }

class Solution {
    public int lengthOfLIS(int[] nums) {
        if (nums == null) return 0;
        int len = nums.length;
        if (len <= 1) return len;


        int[] tail = new int[len];
        tail[0] = nums[0];
        int tailEndIdx = 0;

        for (int i = 1; i < len; i++) {
            if (nums[i] > tail[tailEndIdx]) {
                tailEndIdx++;
                tail[tailEndIdx] = nums[i];
            } else {
                int l = 0;
                int r = tailEndIdx;
                while (l < r) {
                    int mid = l + (r - l) / 2;
                    if (nums[i] > tail[mid]) {
                        l = mid + 1;
                    } else {
                        r = mid;
                    }
                }
                // l will be the first idx which nums[l] >= nums[i]
                tail[l] = nums[i];
            }
        }
        // tailEndIdx is index, the real length of tail is tailEndIdx + 1
        tailEndIdx++;
        return tailEndIdx;

    }
}
// @lc code=end

