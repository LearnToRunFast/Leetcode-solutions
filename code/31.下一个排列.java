/*
 * @lc app=leetcode.cn id=31 lang=java
 *
 * [31] 下一个排列
 *
 * https://leetcode-cn.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (32.02%)
 * Likes:    300
 * Dislikes: 0
 * Total Accepted:    26.9K
 * Total Submissions: 83.9K
 * Testcase Example:  '[1,2,3]'
 *
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * 
 * 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
 * 
 * 必须原地修改，只允许使用额外常数空间。
 * 
 * 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 * 
 */

// @lc code=start
class Solution {
    private static void swap(int[] A, int i, int j) {
        int tmp = A[i];
        A[i] = A[j];
        A[j] = tmp;
    }

    private static void reverse(int[] A, int i, int j) {
        while (i < j)
            swap(A, i++, j--);
    }

    public void nextPermutation(int[] nums) {
        if (nums == null || nums.length <= 1)
            return;
        int tail = nums.length - 2;
        while (tail >= 0 && nums[tail] >= nums[tail + 1])
            tail--;

        if (tail >= 0) {
            int i = nums.length - 1;
            while(nums[i] <= nums[tail]) {
                i--;
            }
            swap(nums, tail, i);
        }
        reverse(nums, tail + 1, nums.length - 1); 
    }
}
// @lc code=end
