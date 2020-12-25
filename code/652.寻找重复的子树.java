import java.util.HashMap;
import java.util.LinkedList;
import java.util.Map;

/*
 * @lc app=leetcode.cn id=652 lang=java
 *
 * [652] 寻找重复的子树
 *
 * https://leetcode-cn.com/problems/find-duplicate-subtrees/description/
 *
 * algorithms
 * Medium (55.03%)
 * Likes:    201
 * Dislikes: 0
 * Total Accepted:    15.2K
 * Total Submissions: 27.5K
 * Testcase Example:  '[1,2,3,4,null,2,4,null,null,4]'
 *
 * 给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
 * 
 * 两棵树重复是指它们具有相同的结构以及相同的结点值。
 * 
 * 示例 1：
 * 
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   2   4
 * ⁠      /
 * ⁠     4
 * 
 * 
 * 下面是两个重复的子树：
 * 
 * ⁠     2
 * ⁠    /
 * ⁠   4
 * 
 * 
 * 和
 * 
 * ⁠   4
 * 
 * 
 * 因此，你需要以列表的形式返回上述重复子树的根结点。
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
// class Solution {
//     private StringBuilder getDuplicates(TreeNode root, Map<String, Integer> count, List<TreeNode> ans) {
//         if (root == null) return new StringBuilder("#");
//         StringBuilder sb = new StringBuilder()
//             .append(root.val)
//             .append(",")
//             .append(getDuplicates(root.left, count, ans))
//             .append(",")
//             .append(getDuplicates(root.right, count, ans));
//         count.put(sb.toString(), count.getOrDefault(sb.toString(), 0) + 1);

//         if (count.get(sb.toString()) == 2) {
//             ans.add(root);
//         }
//         return sb;
//     }
//     public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
//        Map<String, Integer> count = new HashMap<>();
//        List<TreeNode> ans = new LinkedList<>();
//        getDuplicates(root, count, ans);
//        return ans;
//     }
// }
class Solution {
    private int num;
    Map<String, Integer> trees;
    Map<Integer, Integer> count;
    List<TreeNode> ans;

    private int getDuplicates(TreeNode root) {
        if (root == null) return 0;
        // the uid of root is (root.val , x, y) where is x is uid of left tree and y is uid of right tree
        StringBuilder sb = new StringBuilder()
            .append(root.val).append(",")
            .append(getDuplicates(root.left)).append(",")
            .append(getDuplicates(root.right));
        //String serial = root.val + "," + getDuplicates(root.left) + "," + getDuplicates(root.right);
        int uid = trees.computeIfAbsent(sb.toString(), x -> num++);
        count.put(uid, count.getOrDefault(uid, 0) + 1);
        if (count.get(uid) == 2) {
            ans.add(root);
        }
        return uid;
    }
    public List<TreeNode> findDuplicateSubtrees(TreeNode root) {
       num = 1;
       count = new HashMap<>();
       trees = new HashMap<>();
       ans = new LinkedList<>();
       getDuplicates(root);
       return ans;
    }
}
// @lc code=end

