/*
 * @lc app=leetcode.cn id=739 lang=java
 *
 * [739] 每日温度
 *
 * https://leetcode-cn.com/problems/daily-temperatures/description/
 *
 * algorithms
 * Medium (64.46%)
 * Likes:    553
 * Dislikes: 0
 * Total Accepted:    111.6K
 * Total Submissions: 171.4K
 * Testcase Example:  '[73,74,75,71,69,72,76,73]'
 *
 * 请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。如果气温在这之后都不会升高，请在该位置用 0
 * 来代替。
 * 
 * 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4,
 * 2, 1, 1, 0, 0]。
 * 
 * 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
 * 
 */

// @lc code=start
class Solution {
    public int[] dailyTemperatures(int[] T) {
        int[] ans = new int[T.length];
        for (int i = T.length - 2; i >= 0; i--) {
            int curr = i + 1;
            while (T[i] >= T[curr]) {
                if (ans[curr] > 0) {
                    curr += ans[curr];
                } else {
                    curr = i;
                    break;
                }
            }

            ans[i] = curr - i;
        }
        return ans;
    }
}
// class Solution {
//     public int[] dailyTemperatures(int[] T) {
//         int[] ans = new int[T.length];
//         Stack<Integer> stack = new Stack<>();

//         for (int i = 0; i < T.length; i++) {
//             while (!stack.isEmpty() &&  T[i] > T[stack.peek()]) {
//                 int preIdx = stack.pop();
//                 ans[preIdx] = i - preIdx;
//             }
//             stack.push(i);
//         }

//         return ans;
//     }
// }
// @lc code=end

