import java.util.Arrays;
import java.util.LinkedList;
import java.util.Stack;

/*
 * @lc app=leetcode.cn id=84 lang=java
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (41.89%)
 * Likes:    1022
 * Dislikes: 0
 * Total Accepted:    100.6K
 * Total Submissions: 239.8K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 * 
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 * 
 * 
 * 
 * 
 * 
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 * 
 * 
 * 
 * 
 * 
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 * 
 */

// @lc code=start
// class Solution {
//     public int largestRectangleArea(int[] heights) {
//         if (heights.length == 0) return 0;
//         if (heights.length == 1) return heights[0];

//         int ans = 0;
//         int[] newHeights = new int[heights.length + 2];
//         for (int i = 1; i <= heights.length; i++) {
//             newHeights[i] = heights[i - 1];
//         }

//         LinkedList<Integer> stack = new LinkedList<>();
//         stack.push(0); // dummy node, so that no need to check empty stack later.

//         for (int r = 1; r < newHeights.length; r++) {
//             while (newHeights[r] < newHeights[stack.peek()]) {
//                 int idx = stack.poll();
//                 int l = stack.peek();
//                 ans = Math.max(ans, (r - l - 1) * newHeights[idx]);
//             }
//             stack.push(r);
//         }
//         return ans;
//     }
// }
class Solution {
    public int largestRectangleArea(int[] heights) {
        if (heights.length == 0) return 0;
        if (heights.length == 1) return heights[0];

        int ans = 0;

        LinkedList<Integer> stack = new LinkedList<>();
       // stack.push(-1); // dummy node, so that no need to check empty stack later.

        for (int r = 0; r < heights.length; r++) {
            while (!stack.isEmpty() && heights[r] < heights[stack.peek()]) {
                int idx = stack.poll();
                int l = stack.isEmpty() ? -1 : stack.peek();
                ans = Math.max(ans, (r - l - 1) * heights[idx]);
            }
            stack.push(r);
        }
        // clean up
        while (!stack.isEmpty()) {
            int idx = stack.poll();
            int l = stack.isEmpty() ? -1 : stack.peek();
            ans = Math.max(ans, (heights.length - l - 1) * heights[idx]);
        } 
        return ans;
    }
}
// @lc code=end

