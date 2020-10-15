/*
 * @lc app=leetcode.cn id=152 lang=java
 *
 * [152] 乘积最大子数组
 *
 * https://leetcode-cn.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (40.19%)
 * Likes:    792
 * Dislikes: 0
 * Total Accepted:    98.6K
 * Total Submissions: 243.9K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: [2,3,-2,4]
 * 输出: 6
 * 解释: 子数组 [2,3] 有最大乘积 6。
 * 
 * 
 * 示例 2:
 * 
 * 输入: [-2,0,-1]
 * 输出: 0
 * 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 * 
 */

// @lc code=start
// dp version
// class Solution {
//     public int maxProduct(int[] nums) {
//         if (nums == null | nums.length == 0) return -1;
        
//         int prevMin = nums[0];
//         int prevMax = nums[0];
//         int max = nums[0];
//         for (int i = 1; i < nums.length; i++) {
//             int currMax = prevMax;
//             prevMax = Math.max(nums[i],
//                 Math.max(prevMin * nums[i], currMax * nums[i]));

//             prevMin = Math.min(nums[i],
//                 Math.min(prevMin * nums[i], currMax * nums[i]));
//             max = Math.max(max, prevMax);
//         }
//         return max;
//     }
// }

class Solution {
    public int maxProduct(int[] nums) {
        if (nums == null | nums.length == 0) return -1;
        
        int max = nums[0];
        int product = 1;
        for (int i = 0; i < nums.length; i++) {
            product *= nums[i];
            max = Math.max(product, max);
            product = nums[i] == 0 ? 1 : product;
        }
        product = 1;
        for (int i = nums.length - 1; i >= 0; i--) {
            product *= nums[i];
            max = Math.max(product, max);
            product = nums[i] == 0 ? 1 : product;
        }

        return max;

    }
}
// @lc code=end

