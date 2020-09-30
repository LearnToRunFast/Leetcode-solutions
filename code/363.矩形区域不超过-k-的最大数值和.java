/*
 * @lc app=leetcode.cn id=363 lang=java
 *
 * [363] 矩形区域不超过 K 的最大数值和
 *
 * https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/description/
 *
 * algorithms
 * Hard (38.17%)
 * Likes:    124
 * Dislikes: 0
 * Total Accepted:    5.2K
 * Total Submissions: 13.4K
 * Testcase Example:  '[[1,0,1],[0,-2,3]]\n2'
 *
 * 给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。
 * 
 * 示例:
 * 
 * 输入: matrix = [[1,0,1],[0,-2,3]], k = 2
 * 输出: 2 
 * 解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
 * 
 * 
 * 说明：
 * 
 * 
 * 矩阵内的矩形区域面积必须大于 0。
 * 如果行数远大于列数，你将如何解答呢？
 * 
 * 
 */

// @lc code=start
class Solution {
    private int findMax(int[] dp, int k) {
        int max = dp[0], curr = dp[0];

        for(int i = 1; i < dp.length; i++) {
            curr = Integer.max(dp[i], curr + dp[i]);
            max = Integer.max(max, curr);
        }
        if (max <= k) return max;

        
        max = Integer.MIN_VALUE;
        for(int i = 0; i < dp.length; i++) {
            int sum = 0;
            for(int j = i; j < dp.length; j++) {
                sum += dp[j];
                
                if (sum <= k && sum > max) {
                    max = sum;
                }
                if (sum == k) return sum;
            }
        }
        return max;
    }
    public int maxSumSubmatrix(int[][] matrix, int k) {

        int rows = matrix.length, cols = matrix[0].length;
        int max = Integer.MIN_VALUE;

        for(int c = 0; c < cols; c++) {
            // start a new dp when new column appear.
            int[] dp = new int[rows];
            for(int i = c; i < cols; i++) {
                for(int r = 0; r < rows; r++) {
                    dp[r] += matrix[r][i];
                }
                max = Integer.max(max, findMax(dp, k));
                if (max == k) return k;
            }
        }

        return max;
    }
}
// @lc code=end

