/*
 * @lc app=leetcode.cn id=685 lang=java
 *
 * [685] 冗余连接 II
 *
 * https://leetcode-cn.com/problems/redundant-connection-ii/description/
 *
 * algorithms
 * Hard (44.66%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 36.2K
 * Testcase Example:  '[[1,2],[1,3],[2,3]]'
 *
 * 在本问题中，有根树指满足以下条件的有向图。该树只有一个根节点，所有其他节点都是该根节点的后继。每一个节点只有一个父节点，除了根节点没有父节点。
 * 
 * 输入一个有向图，该图由一个有着N个节点 (节点值不重复1, 2, ..., N)
 * 的树及一条附加的边构成。附加的边的两个顶点包含在1到N中间，这条附加的边不属于树中已存在的边。
 * 
 * 结果图是一个以边组成的二维数组。 每一个边 的元素是一对 [u, v]，用以表示有向图中连接顶点 u 和顶点 v 的边，其中 u 是 v
 * 的一个父节点。
 * 
 * 返回一条能删除的边，使得剩下的图是有N个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。
 * 
 * 示例 1:
 * 
 * 输入: [[1,2], [1,3], [2,3]]
 * 输出: [2,3]
 * 解释: 给定的有向图如下:
 * ⁠ 1
 * ⁠/ \
 * v   v
 * 2-->3
 * 
 * 
 * 示例 2:
 * 
 * 输入: [[1,2], [2,3], [3,4], [4,1], [1,5]]
 * 输出: [4,1]
 * 解释: 给定的有向图如下:
 * 5 <- 1 -> 2
 * ⁠    ^    |
 * ⁠    |    v
 * ⁠    4 <- 3
 * 
 * 
 * 注意:
 * 
 * 
 * 二维数组大小的在3到1000范围内。
 * 二维数组中的每个整数在1到N之间，其中 N 是二维数组的大小。
 * 
 * 
 */

// @lc code=start
class Solution {
    int[] parent;
    int[] indegree;
    private int root(int p) {
        // not root
        while (parent[p] != p) {
            parent[p] = parent[parent[p]]; // path compression
            p = parent[p];
        }
        return p;
    }
    // return false if it's already connect, otherwise true.
    private boolean union(int p, int q) {
        int rootP = root(p);
        int rootQ = root(q);
        if (rootP == rootQ) return false;

        parent[rootP] = rootQ;
        return true;
    }

    public int[] findRedundantDirectedConnection(int[][] edges) {
        // either one node has two indegree or there is a directed cycle.
        // if there is not two indegree node, it must contains cycle
        // if two indegree
        //   1. contains cycle
        //   2. does not contain cycle
        int[] ans = new int[2];
        parent = new int[edges.length + 1];
        for (int i = 0; i < parent.length; i++) {
            parent[i] = i;
        }

        int[][] twoIndegreeEdges = new int[2][2];
        indegree = new int[parent.length];
        int twoIndegreeVertex = 0;
        for (int[] edge: edges) {
            // u -> v
            indegree[edge[1]]++;

            // two indegree
            if (indegree[edge[1]] == 2) {
                twoIndegreeVertex = edge[1];
                twoIndegreeEdges[1] = edge;
            } else {
                if (!union(edge[0], edge[1])) {
                    ans = edge;
                }
            }
        }

        if (twoIndegreeVertex != 0) {
            for (int[] edge : edges) {
                if (edge[1] == twoIndegreeVertex) {
                    twoIndegreeEdges[0] = edge;
                    break;
                }
            }
            // if all the root are the same,
            // means it's a tree
            // so we return extra edge that does not included in the tree.
            int commonRoot = root(1);
            for (int i = 2; i < parent.length; i++) {
                // if it contains cycle or it has diff root
                if (ans[0] != 0 || commonRoot != root(i)) {
                    return twoIndegreeEdges[0];
                }
            }
            return twoIndegreeEdges[1];
        }

        return ans;
    }
}
// @lc code=end

