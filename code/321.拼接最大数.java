/*
 * @lc app=leetcode.cn id=321 lang=java
 *
 * [321] 拼接最大数
 *
 * https://leetcode-cn.com/problems/create-maximum-number/description/
 *
 * algorithms
 * Hard (43.05%)
 * Likes:    344
 * Dislikes: 0
 * Total Accepted:    22.2K
 * Total Submissions: 51.5K
 * Testcase Example:  '[3,4,6,5]\n[9,1,2,5,8,3]\n5'
 *
 * 给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n)
 * 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
 * 
 * 求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
 * 
 * 说明: 请尽可能地优化你算法的时间和空间复杂度。
 * 
 * 示例 1:
 * 
 * 输入:
 * nums1 = [3, 4, 6, 5]
 * nums2 = [9, 1, 2, 5, 8, 3]
 * k = 5
 * 输出:
 * [9, 8, 6, 5, 3]
 * 
 * 示例 2:
 * 
 * 输入:
 * nums1 = [6, 7]
 * nums2 = [6, 0, 4]
 * k = 5
 * 输出:
 * [6, 7, 6, 0, 4]
 * 
 * 示例 3:
 * 
 * 输入:
 * nums1 = [3, 9]
 * nums2 = [8, 9]
 * k = 3
 * 输出:
 * [9, 8, 9]
 * 
 */

// @lc code=start
class Solution {
    private void printArr(int[] a) {
        if (a.length == 0) return;
        StringBuilder sb = new StringBuilder();
        sb.append("[");
        for (int num : a) {
            sb.append(num).append(",");
        }
        sb.replace(sb.length() - 1, sb.length(), "]");
        System.out.println(sb.toString());
    }
    private int[] merge(int[] a, int[] b) {
        int[] seq = new int[a.length + b.length];
        int p1 = 0, p2 = 0, l1 = 0, l2 = 0;
        int curr = 0;
        while (l1 < a.length && l2 < b.length) {
            if (a[l1] > b[l2]) {
                seq[curr++] = a[l1++];
                continue;
            }
            if (a[l1] < b[l2]) {
                seq[curr++] = b[l2++];
                continue;
            } 
            // a[l1] == b[l2]
            p1 = l1 + 1;
            p2 = l2 + 1;
            while (p1 < a.length && p2 < b.length && a[p1] == b[p2]) {
                p1++;
                p2++;
            }
            int ans1 = p1 == a.length ? Integer.MIN_VALUE : a[p1];
            int ans2 = p2 == b.length ? Integer.MIN_VALUE : b[p2];
            if (ans1 >= ans2) {
                seq[curr++] = a[l1++];
            } else {
                seq[curr++] = b[l2++];
            }
        }
        while (l1 < a.length) {
            seq[curr++] = a[l1++];
        }
        while (l2 < b.length) {
            seq[curr++] = b[l2++];
        }
        return seq;
    }
    private int[] getMaxSequence(int[] nums, int k) {
        if (k == 0) return new int[k];
        if (nums.length == k) return nums;

        int[] maxSeq = new int[k];

        int extra = nums.length - k;
        int curr = -1;

        for (int num : nums) {
            if (extra == 0) {
                maxSeq[++curr] = num;
                continue;
            }
            while (curr >= 0 && maxSeq[curr] < num && extra > 0) {
                extra--;
                curr--;
            }
            if (curr >= k - 1) {
                extra--;
            } else {
                maxSeq[++curr] = num;
            }
        }
        return maxSeq;
    }
    private boolean less(int[] a, int[] b) {
        for (int i = 0; i < a.length; i++) {
            if (a[i] > b[i]) {
                return false;
            } else if (a[i] < b[i]) {
                return true;
            }
        }
        return false;
    }
    public int[] maxNumber(int[] nums1, int[] nums2, int k) {
        if (nums1.length > nums2.length) {
            return maxNumber(nums2, nums1, k);
        }
        int[] maxSeq = new int[k];
        int i = Math.min(nums1.length, k);
        int end = Math.max(k - nums2.length, 0);
        for(; i >= end; i--) { 
            //System.out.println("==================");
            int[] a = getMaxSequence(nums1, i);
            //printArr(a);
            int[] b = getMaxSequence(nums2, k - i);
            //printArr(b);
            int[] combined = merge(a, b);
            // printArr(combined);
            if (less(maxSeq, combined)) {
                maxSeq = combined;
            }
        }
        return maxSeq;
    }
}
// @lc code=end

