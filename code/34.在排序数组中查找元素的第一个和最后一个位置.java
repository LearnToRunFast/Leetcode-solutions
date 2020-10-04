/*
 * @lc app=leetcode.cn id=34 lang=java
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (40.22%)
 * Likes:    569
 * Dislikes: 0
 * Total Accepted:    125.1K
 * Total Submissions: 311K
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 如果数组中不存在目标值，返回 [-1, -1]。
 * 
 * 示例 1:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 8
 * 输出: [3,4]
 * 
 * 示例 2:
 * 
 * 输入: nums = [5,7,7,8,8,10], target = 6
 * 输出: [-1,-1]
 * 
 */

// @lc code=start
class Solution {
    // this version binary search will find first element if there are many of them
    // if there is not such element, it will return the index that such element belong to.
    private int binarySearch(int[] nums, int target) {
        int l = 0, r = nums.length;

        while(l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] >= target) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        return l;
    }

    public int[] searchRange(int[] nums, int target) {
        int i = binarySearch(nums, target);
        int j = binarySearch(nums, target + 1);
        // if cant find it 
        if (i == nums.length || nums[i] != target) return new int[]{-1, -1};

        return new int[]{i, j - 1};


    }
}
// @lc code=end

