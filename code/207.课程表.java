import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Queue;

/*
 * @lc app=leetcode.cn id=207 lang=java
 *
 * [207] 课程表
 *
 * https://leetcode-cn.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (54.06%)
 * Likes:    603
 * Dislikes: 0
 * Total Accepted:    79.4K
 * Total Submissions: 145.9K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
 * 
 * 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
 * 
 * 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: 2, [[1,0]] 
 * 输出: true
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
 * 
 * 示例 2:
 * 
 * 输入: 2, [[1,0],[0,1]]
 * 输出: false
 * 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
 * 你可以假定输入的先决条件中没有重复的边。
 * 1 <= numCourses <= 10^5
 * 
 * 
 */

// @lc code=start
// // // BSF
// class Solution {
//     public boolean canFinish(int numCourses, int[][] prerequisites) {
//         // idx is course number, value is inDegree of that course number
//         int[] inDegrees = new int[numCourses];
//         Map<Integer, ArrayList<Integer>> adj = new HashMap<>();
//         for (int i = 0; i < numCourses; i++) {
//             adj.put(i, new ArrayList<>());
//         }

        
//         for (int[] pre : prerequisites) {
//             // increase indegress of course number in pre[0]
//             inDegrees[pre[0]]++; 

//             // add to adj matrix
//             adj.get(pre[1]).add(pre[0]);
//         }

//         // queue for BFS
//         Queue<Integer> queue = new LinkedList<>();
//         // add indegree of 0 into queue
//         for (int i = 0; i < numCourses; i++) {
//             if (inDegrees[i] == 0) {
//                 queue.add(i);
//             }
//         }
//         // start BFS
//         while (!queue.isEmpty()) {
//             int currCourse = queue.poll();
//             numCourses--;

//             // look at course number adj to current course number
//             for(int adjCourse : adj.get(currCourse)) {
//                 // decrease it first as current course is been removed
//                 if (--inDegrees[adjCourse] == 0) {
//                     queue.add(adjCourse);
//                 }
//             }
//         }

//         return numCourses == 0;
        
//     }
// }

class Solution {
    private boolean dfs(List<List<Integer>> adj, int[] visited, int i) {
        if (visited[i] == 1) return false; //already visited, there is a cycle
        if (visited[i] == -1) return true; // its been visited by other dfs.

        visited[i] = 1;
        for(int neighbour : adj.get(i)) {
            if (!dfs(adj, visited, neighbour)) return false;
        }
        visited[i] = -1;
        return true;
    }
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        List<List<Integer>> adj = new ArrayList<>();
        for(int i = 0; i < numCourses; i++) {
            adj.add(new ArrayList<>());
        }
        int[] visited = new int[numCourses];

        for(int[] pre : prerequisites) {
            adj.get(pre[1]).add(pre[0]);
        }
        for(int i = 0; i < numCourses; i++) {
            if (!dfs(adj, visited, i)) return false;
        }
        return true;
    }
}
// @lc code=end

