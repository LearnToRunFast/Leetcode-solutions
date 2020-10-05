/*
 * @lc app=leetcode.cn id=81 lang=java
 *
 * [81] 搜索旋转排序数组 II
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (35.78%)
 * Likes:    228
 * Dislikes: 0
 * Total Accepted:    41.3K
 * Total Submissions: 114.9K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
 * 
 * 编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
 * 
 * 示例 1:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 0
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [2,5,6,0,0,1,2], target = 3
 * 输出: false
 * 
 * 进阶:
 * 
 * 
 * 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
 * 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 * 
 * 
 */

// @lc code=start
class Solution {
    private boolean search(int[] nums, int l, int r, int target) {
        if (r == -1) return false;

        while (l < r) {
            int mid = l + (r - l) / 2;
            if (target == nums[mid]) return true;

            // no sorted
            if (nums[l] > nums[mid]) {
                if (target > nums[mid] && target <= nums[r]) {
                    l = mid + 1;
                } else {
                    r = mid;
                }

            // sorted
            } else if (nums[l] < nums[mid]) {
                if (target > nums[mid] || target < nums[l]) {
                    l = mid + 1;
                } else {
                    r = mid;
                }

            } else {
                // if nums[l] == nums[mid] == nums[r]
                if (nums[mid] == nums[r]) {
                    return search(nums, l + 1, mid - 1, target) || 
                           search(nums, mid + 1, r - 1, target);
                }

                if ((nums[r] < nums[l] && target > nums[l]) ||
                    (nums[r] > nums[l] && target < nums[l])) {
                    return false;
                }

                l = mid + 1;
            }
        }

        return nums[l] == target ? true : false;
        
    }
    public boolean search(int[] nums, int target) {

        return search(nums, 0, nums.length - 1, target);
    }
}
// @lc code=end

