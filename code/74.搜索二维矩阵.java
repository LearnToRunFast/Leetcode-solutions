/*
 * @lc app=leetcode.cn id=74 lang=java
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.85%)
 * Likes:    89
 * Dislikes: 0
 * Total Accepted:    18.3K
 * Total Submissions: 51.1K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 * 
 * 
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 * 
 */

// @lc code=start
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0)
            return false;
        int i = 0, j = matrix[i].length - 1;
        int n = matrix.length, m = j;

        while (i < n && j >= 0) {
            if (matrix[i][j] < target) {
                i++;
            } else {
                if (matrix[i][j] == target)
                    return true;
                int l = 0, h = m;
                while (l < h) {
                    int mid = (l + h) >>> 1;
                    if (matrix[i][mid] < target) {
                        l = mid + 1;
                    } else {
                        h = mid;
                    }
                }
                return matrix[i][l] == target;
            }

        }

        return false;
    }
}
// @lc code=end
