/*
 * @lc app=leetcode.cn id=283 lang=java
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (62.11%)
 * Likes:    780
 * Dislikes: 0
 * Total Accepted:    225.6K
 * Total Submissions: 360.4K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * 
 * 示例:
 * 
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 * 
 * 说明:
 * 
 * 
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 * 
 * 
 */

// @lc code=start
class Solution {

    private int nextZeroIdx(int[] nums, int i) {
        while (i < nums.length && nums[i] != 0) {
            i++;
        }
        return i == nums.length ? -1 : i;
    }

    public void moveZeroes(int[] nums) {
        if (nums == null || nums.length == 0) return;

        int zeroIdx = nextZeroIdx(nums, 0);
        if (zeroIdx == -1) return;

        for (int i = zeroIdx + 1; i < nums.length; i++) {
            if (nums[i] != 0) {
                nums[zeroIdx] = nums[i];
                nums[i] = 0;
                zeroIdx = nextZeroIdx(nums, zeroIdx + 1);
                if (zeroIdx == -1) return;
            }

        }
    }
}
// @lc code=end

