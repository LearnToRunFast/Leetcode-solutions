/*
 * @lc app=leetcode.cn id=35 lang=java
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (46.65%)
 * Likes:    702
 * Dislikes: 0
 * Total Accepted:    256.3K
 * Total Submissions: 549K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 * 
 * 你可以假设数组中无重复元素。
 * 
 * 示例 1:
 * 
 * 输入: [1,3,5,6], 5
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 输入: [1,3,5,6], 2
 * 输出: 1
 * 
 * 
 * 示例 3:
 * 
 * 输入: [1,3,5,6], 7
 * 输出: 4
 * 
 * 
 * 示例 4:
 * 
 * 输入: [1,3,5,6], 0
 * 输出: 0
 * 
 * 
 */

// @lc code=start
class Solution {

    public int searchInsert(int[] nums, int target) {
        int l = 0, r = nums.length;

        // [l, r)
        while(l < r) {
            int mid = l + ( r - l) / 2;

            if (nums[mid] == target) return mid;
            
            // since r is not included, so here should put mid instead of mid + 1
            if (nums[mid] > target) {
                r = mid;
            } else {
            // nums[mid] < target
            // since l is included, so l = mid + 1 as mid is not the index we want.
                l = mid + 1;
            }
        }
        return l;

    }
}
// @lc code=end

