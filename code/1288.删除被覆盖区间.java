import java.lang.reflect.Array;
import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=1288 lang=java
 *
 * [1288] 删除被覆盖区间
 *
 * https://leetcode-cn.com/problems/remove-covered-intervals/description/
 *
 * algorithms
 * Medium (55.46%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    6.8K
 * Total Submissions: 12.3K
 * Testcase Example:  '[[1,4],[3,6],[2,8]]'
 *
 * 给你一个区间列表，请你删除列表中被其他区间所覆盖的区间。
 * 
 * 只有当 c <= a 且 b <= d 时，我们才认为区间 [a,b) 被区间 [c,d) 覆盖。
 * 
 * 在完成所有删除操作后，请你返回列表中剩余区间的数目。
 * 
 * 
 * 
 * 示例：
 * 
 * 
 * 输入：intervals = [[1,4],[3,6],[2,8]]
 * 输出：2
 * 解释：区间 [3,6] 被区间 [2,8] 覆盖，所以它被删除了。
 * 
 * 
 * 
 * 
 * 提示：​​​​​​
 * 
 * 
 * 1 <= intervals.length <= 1000
 * 0 <= intervals[i][0] < intervals[i][1] <= 10^5
 * 对于所有的 i != j：intervals[i] != intervals[j]
 * 
 * 
 */

// @lc code=start
// class Solution {
//     public int removeCoveredIntervals(int[][] intervals) {
//         Arrays.sort(intervals, (int[] a, int[] b) -> {
//             return a[0] == b[0] ? b[1] - a[1] : a[0] - b[0];
//         });

//         int right = intervals[0][1];
//         int ans = 0;
//         for (int i = 1; i <intervals.length; i++) {
//             int[] curr = intervals[i];
//             // completely smaller
//             if (right >= curr[1]) {
//                 ans++;
//             }
//             // partial overlap
//             else if (curr[0] <= right && right < curr[1]) {
//                 right = curr[1];
//             }
//             // disjoin
//             else if (right < curr[0]) {
//                 right = curr[1];
//             }
//         }
//         return intervals.length - ans;
//     }
// }
class Solution {
    public int removeCoveredIntervals(int[][] intervals) {
        Arrays.sort(intervals, (int[] a, int[] b) -> {
            return a[0] == b[0] ? b[1] - a[1] : a[0] - b[0];
        });

        int ans = 0;
        int preEnd = 0, currEnd = 0;
        for (int[] interval: intervals) {
            currEnd = interval[1];

            if (currEnd > preEnd) {
                ans++;
                preEnd = currEnd;
            }
        }
        return ans;
    }
}
// @lc code=end

