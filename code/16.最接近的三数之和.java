import javax.naming.spi.DirStateFactory.Result;

/*
 * @lc app=leetcode.cn id=16 lang=java
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.65%)
 * Likes:    272
 * Dislikes: 0
 * Total Accepted:    47.6K
 * Total Submissions: 114.1K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 * 
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 * 
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 * 
 * 
 */

// @lc code=start
class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int res = nums[0] + nums[1] + nums[2];
        for (int i = 0; i < nums.length - 2; i++) {
            int left = i + 1;
            int right = nums.length - 1;
            while (left != right) {

                int sum = nums[i] + nums[left] + nums[right];
                if (Math.abs(sum - target) < Math.abs(res - target)) {
                    res = sum;
                }
                if (sum > target) {
                    // for duplicate value ahead
                    while (left != right && nums[right] == nums[--right])
                        ;
                } else {
                    if (sum == target)
                        return sum;
                    while (left != right && nums[left] == nums[++left])
                        ;

                }
            }
            while (i < nums.length - 2 && nums[i] == nums[i + 1])
                i++;
        }
        return res;
    }
}
// @lc code=end
