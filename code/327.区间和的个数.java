/*
 * @lc app=leetcode.cn id=327 lang=java
 *
 * [327] 区间和的个数
 *
 * https://leetcode-cn.com/problems/count-of-range-sum/description/
 *
 * algorithms
 * Hard (35.65%)
 * Likes:    245
 * Dislikes: 0
 * Total Accepted:    17.5K
 * Total Submissions: 41.4K
 * Testcase Example:  '[-2,5,-1]\n-2\n2'
 *
 * 给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
 * 区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。
 * 
 * 说明:
 * 最直观的算法复杂度是 O(n^2) ，请在此基础上优化你的算法。
 * 
 * 示例:
 * 
 * 输入: nums = [-2,5,-1], lower = -2, upper = 2,
 * 输出: 3 
 * 解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。
 * 
 * 
 */

// @lc code=start
class Solution {

    private int midCount(long[] sum, int lower, int upper, int left, int mid, int right) {
            int i = left;
            int l = mid + 1;
            int count = 0;
            while (i <= mid) {
                while (l <= right && sum[l] - sum[i] < lower) {
                    l++;
                }
                int r = l;
                while (r <= right && sum[r] - sum[i] <= upper) {
                    r++;
                }
                count += r - l;
                i++;
            }
            return count;
    }
    private int countRangeSum(long[] sum, int lower, int upper, int l, int r) {
        if (l == r) return 0;

        int mid = l + (r - l) / 2;
        int leftCount = countRangeSum(sum, lower, upper, l, mid);
        int rightCount = countRangeSum(sum, lower, upper, mid + 1, r);
        int count = leftCount + rightCount;

        count += midCount(sum, lower, upper, l, mid, r);

        long [] sorted = new long[r - l + 1];
        int left = l;
        int right = mid + 1;
        int i = 0;
        while (left <= mid && right <= r) {
            if (sum[left] > sum[right]) {
                sorted[i++] = sum[right++];
            } else {
                sorted[i++] = sum[left++];
            }
        }
        while (left <= mid) {
            sorted[i++] = sum[left++];
        }
        while (right <= r) {
            sorted[i++] = sum[right++];
        }

        for (int k = 0; k < sorted.length; k++) {
            sum[k + l] = sorted[k];
        }
        return count;
    }
    public int countRangeSum(int[] nums, int lower, int upper) {
        long[] sum = new long[nums.length + 1];
        for (int i = 0; i < nums.length; i++) {
            sum[i + 1] = nums[i] + sum[i];
        }
        return countRangeSum(sum, lower, upper, 0, sum.length - 1);

    }
}


// @lc code=end

