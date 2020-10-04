/*
 * @lc app=leetcode.cn id=33 lang=java
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (36.11%)
 * Likes:    908
 * Dislikes: 0
 * Total Accepted:    162.5K
 * Total Submissions: 420.7K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
 * 
 * ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
 * 
 * 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 * 
 * 你可以假设数组中不存在重复的元素。
 * 
 * 你的算法时间复杂度必须是 O(log n) 级别。
 * 
 * 示例 1:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 0
 * 输出: 4
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [4,5,6,7,0,1,2], target = 3
 * 输出: -1
 * 
 */

// @lc code=start
class Solution {
    public int search(int[] nums, int target) {
        //put -1 here to ensure r will never go to n
        // so in the last return no need to check l != nums.length
        int l = 0, r = nums.length - 1; 

        while(l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] == target) return mid;

            // three conditions on left
            //   nums[0] <= target <= nums[i]
            //              target <= nums[i] < nums[0]
            //                        nums[i] < nums[0] <= target
            if ((nums[0] <= target) ^ (target <= nums[mid]) ^ (nums[mid] < nums[0])) {
                    l = mid + 1;
                } else {
                    r = mid;
                }

        }
    
        return nums[l] == target ? l : -1;
    }
}
// @lc code=end

