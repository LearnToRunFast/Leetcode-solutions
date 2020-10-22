import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=239 lang=java
 *
 * [239] 滑动窗口最大值
 *
 * https://leetcode-cn.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (48.86%)
 * Likes:    598
 * Dislikes: 0
 * Total Accepted:    82.4K
 * Total Submissions: 167.5K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k
 * 个数字。滑动窗口每次只向右移动一位。
 * 
 * 返回滑动窗口中的最大值。
 * 
 * 
 * 
 * 进阶：
 * 
 * 你能在线性时间复杂度内解决此题吗？
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
 * 输出: [3,3,5,5,6,7] 
 * 解释: 
 * 
 * ⁠ 滑动窗口的位置                最大值
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7       3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * 1 <= k <= nums.length
 * 
 * 
 */

// @lc code=start
// double ended queue (dequeue)
// class Solution {

//     public int[] maxSlidingWindow(int[] nums, int k) {
//         if (nums == null || nums.length < 2 || k == 1) return nums;

//         int ans[] = new int[nums.length - k + 1];
//         LinkedList<Integer> queue = new LinkedList<>();

//         for (int i = 0; i < nums.length; i++) {
//             // makre sure queue is in decrease order
//             while (!queue.isEmpty() && nums[queue.peekLast()] <= nums[i]) {
//                 queue.pollLast();
//             }
//             // add index instead of value.
//             queue.add(i);

//             // check range
//             if (queue.peek() <= i - k) {
//                 queue.poll();
//             }

//             // check completeness of window
//             int idx = i + 1 - k;
//             if (idx >= 0) {
//                 ans[idx] = nums[queue.peek()];
//             }
//         }
//         return ans;
//     }
// }

class Solution {

    public int[] maxSlidingWindow(int[] nums, int k) {
        if (nums == null || nums.length < 2 || k == 1) return nums;
        int n = nums.length;
        int[] left = new int[n];
        int[] right = new int[n];
        left[0] = nums[0];
        right[n - 1] = nums[n - 1];
        for (int i = 1, j = n - 2; i < n; i++, j--) {
            if (i % k == 0) {
                left[i] = nums[i];
            } else {
                left[i] = Math.max(left[i - 1], nums[i]);
            }

            if ((j + 1) % k == 0) {
                right[j] = nums[j];
            } else {
                right[j] = Math.max(right[j + 1], nums[j]);
            }
        }

        int[] ans = new int[n - k + 1];
        for (int i = 0; i < n - k + 1; i++) {
            ans[i] = Math.max(left[i + k - 1], right[i]);
        }
        return ans;
    }
}
// @lc code=end

