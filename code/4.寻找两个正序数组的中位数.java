/*
 * @lc app=leetcode.cn id=4 lang=java
 *
 * [4] 寻找两个正序数组的中位数
 *
 * https://leetcode-cn.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (38.64%)
 * Likes:    3246
 * Dislikes: 0
 * Total Accepted:    265.5K
 * Total Submissions: 684.3K
 * Testcase Example:  '[1,3]\n[2]'
 *
 * 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
 * 
 * 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums1 = [1,3], nums2 = [2]
 * 输出：2.00000
 * 解释：合并数组 = [1,2,3] ，中位数 2
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums1 = [1,2], nums2 = [3,4]
 * 输出：2.50000
 * 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
 * 
 * 
 * 示例 3：
 * 
 * 输入：nums1 = [0,0], nums2 = [0,0]
 * 输出：0.00000
 * 
 * 
 * 示例 4：
 * 
 * 输入：nums1 = [], nums2 = [1]
 * 输出：1.00000
 * 
 * 
 * 示例 5：
 * 
 * 输入：nums1 = [2], nums2 = []
 * 输出：2.00000
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 * 
 * 
 */

// @lc code=start
class Solution {

    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        // the left size (m + n + 1) / 2
        // note that if it's even (m + n) / 2 == (m + n + 1) / 2
        // if it's odd (m + n + 1) / 2

        // draw a line, the left must be <= to right
        // by make sure the upper left is smaller than lower right
        // and the upper right must be greater than lower left

        //making sure len of nums1 is <= nums2
        if (nums1.length > nums2.length) {
            return findMedianSortedArrays(nums2, nums1);
        }

        int m = nums1.length, n = nums2.length;
        // to prevent overflow of (m + n + 1) / 2
        int totalLeft = m + (n - m + 1) / 2;

        int l = 0, r = m;

        while(l < r) {
            int i = l + (r - l) / 2; //add 1 here to prevent infinite loop
            int j = totalLeft - i;
            // nums1[i - 1] <= nums2[j] && nums2[j - 1] <= nums1[i]
            if (nums2[j - 1] > nums1[i]) {
                l = i + 1;
            } else {
                r = i;
            }
        }

        int i = l;
        int j = totalLeft - i;
        int nums1LeftMax = i == 0 ? Integer.MIN_VALUE : nums1[i - 1];
        int nums1RightMin = i == m ? Integer.MAX_VALUE : nums1[i];
        int nums2LeftMax = j == 0 ? Integer.MIN_VALUE : nums2[j - 1];
        int num2RightMin = j == n ? Integer.MAX_VALUE : nums2[j];

        if ((m + n) % 2 == 1) {
            return Math.max(nums1LeftMax, nums2LeftMax);
        }

        return (Math.max(nums1LeftMax, nums2LeftMax) + 
                Math.min(nums1RightMin, num2RightMin)) / 2.0;
    }
}
// @lc code=end

