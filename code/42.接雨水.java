import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=42 lang=java
 *
 * [42] 接雨水
 *
 * https://leetcode-cn.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (53.04%)
 * Likes:    1802
 * Dislikes: 0
 * Total Accepted:    163.6K
 * Total Submissions: 308.5K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 * 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：height = [4,2,0,3,2,5]
 * 输出：9
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == height.length
 * 0 
 * 0 
 * 
 * 
 */

// @lc code=start
// dp
// class Solution {
//     public int trap(int[] height) {
//         if (height == null || height.length < 3) return 0;
//         int n = height.length;
//         int[] leftMax = new int[n];
//         leftMax[0] = height[0];
//         for (int i = 1; i < n; i++) {
//             leftMax[i] = Math.max(leftMax[i - 1], height[i]);
//         }

//         int[] rightMax = new int[n];
//         rightMax[n - 1] = height[n - 1];
//         for (int i = n - 2; i >= 0; i--) {
//             rightMax[i] = Math.max(rightMax[i + 1], height[i]);
//         }

//         int ans = 0;
//         for (int i = 0; i < n; i++) {
//             ans += Math.min(leftMax[i], rightMax[i]) - height[i];
//         }
//         return ans;
//     }
// }
// Stack
// class Solution {
//     public int trap(int[] height) {
//         if (height == null || height.length < 3) return 0;
//         int n = height.length;
//         LinkedList<Integer> stack = new LinkedList<>();;
//         int ans = 0;
//         for (int i = 0; i < n; i++) {
//             while (!stack.isEmpty() && height[i] > height[stack.peek()]) {
//                 int prevIdx = stack.pop();
//                 if (stack.isEmpty()) break;
//                 int distance = i - stack.peek() - 1;
//                 int h = Math.min(height[stack.peek()], height[i]) - height[prevIdx];
//                 ans += distance * h;
//             }
//             stack.push(i);
//         }
//         return ans;
//     }
// }
class Solution {
    public int trap(int[] height) {
        if (height == null || height.length < 3) return 0;
        int n = height.length;
        int ans = 0;
        int l = 0;
        int r = n - 1;
        int lMax = 0;
        int rMax = 0;
        while (l < r) {
            if (height[l] < height[r]) {
                if (height[l] > lMax) {
                    lMax = height[l];
                } else {
                    ans += lMax - height[l];
                }
                l++;
            } else {
                if (height[r] > rMax) {
                    rMax = height[r];
                } else {
                    ans += rMax - height[r];
                }
                r--;
            }
        }
        return ans;
    }
}
// @lc code=end

