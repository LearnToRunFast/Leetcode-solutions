import java.util.ArrayList;
import java.util.Deque;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=863 lang=java
 *
 * [863] 二叉树中所有距离为 K 的结点
 *
 * https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/description/
 *
 * algorithms
 * Medium (53.08%)
 * Likes:    206
 * Dislikes: 0
 * Total Accepted:    9.7K
 * Total Submissions: 18.2K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n2'
 *
 * 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。
 * 
 * 返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
 * 输出：[7,4,1]
 * 解释：
 * 所求结点为与目标结点（值为 5）距离为 2 的结点，
 * 值分别为 7，4，以及 1
 * 
 * 
 * 
 * 注意，输入的 "root" 和 "target" 实际上是树上的结点。
 * 上面的输入仅仅是对这些对象进行了序列化描述。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 给定的树是非空的。
 * 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
 * 目标结点 target 是树上的结点。
 * 0 <= K <= 1000.
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
// class Solution {
//     private Map<TreeNode, TreeNode> parents;
//     private Set<TreeNode> visited;
//     private List<Integer> ans;
//     private void getParent(TreeNode root, TreeNode parent) {
//         if (root == null) return;
//         parents.put(root, parent);
//         getParent(root.left, root);
//         getParent(root.right, root);
//     }
//     public List<Integer> distanceK(TreeNode root, TreeNode target, int K) {
//         parents = new HashMap<>();
//         getParent(root, null);
//         visited = new HashSet<>();
//         ans = new ArrayList<>();
//         Deque<TreeNode> queue = new LinkedList<>();
//         queue.add(target);
//         while (!queue.isEmpty()) {
//             if (K-- == 0) {
//                 while (!queue.isEmpty()) {
//                     ans.add(queue.pop().val);
//                 }
//                 continue;
//             }
//             int size = queue.size();
//             for (int i = 0; i < size; i++) {
//                 TreeNode node = queue.poll();
//                 visited.add(node);
//                 if (node.left != null && !visited.contains(node.left)) {
//                     queue.add(node.left);
//                 }
//                 if (node.right != null && !visited.contains(node.right)) {
//                     queue.add(node.right);
//                 }
//                 TreeNode p = parents.get(node);
//                 if (p != null && !visited.contains(p)) {
//                     queue.add(p);
//                 }
//             }
//         }
//         return ans;
//     }
// }
class Solution {

    private List<Integer> ans;
    private int K;
    private TreeNode target;
    private void addNodes(TreeNode root, int dist) {
        if (dist > K || root == null) return;
        if (dist == K) {
            ans.add(root.val);
            return;
        }
        addNodes(root.left, dist + 1);
        addNodes(root.right, dist + 1);
    }
    private int dfs(TreeNode root) {
        if (root == null) return -1;
        if (root == target) {
            addNodes(root, 0);
            return 1;
        }

        int left = dfs(root.left);
        int right = dfs(root.right);

        if (left != -1) {
            if (left == K) {
                ans.add(root.val);
            }
            addNodes(root.right, left + 1);
            return left + 1;
        }
        if (right != -1) {
            if (right == K) {
                ans.add(root.val);
            }
            addNodes(root.left, right + 1);
            
            return right + 1;
        }
        return -1;
    }
    public List<Integer> distanceK(TreeNode root, TreeNode target, int K) {
        ans = new ArrayList<>();
        this.K = K;
        this.target = target;
        dfs(root);
        return ans;
    }
}
// @lc code=end

