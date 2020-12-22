import java.util.Deque;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.Set;

import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=606 lang=java
 *
 * [606] 根据二叉树创建字符串
 *
 * https://leetcode-cn.com/problems/construct-string-from-binary-tree/description/
 *
 * algorithms
 * Easy (55.38%)
 * Likes:    168
 * Dislikes: 0
 * Total Accepted:    19.6K
 * Total Submissions: 35.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
 * 
 * 空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 二叉树: [1,2,3,4]
 * ⁠      1
 * ⁠    /   \
 * ⁠   2     3
 * ⁠  /    
 * ⁠ 4     
 * 
 * 输出: "1(2(4))(3)"
 * 
 * 解释: 原本将是“1(2(4)())(3())”，
 * 在你省略所有不必要的空括号对之后，
 * 它将是“1(2(4))(3)”。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: 二叉树: [1,2,3,null,4]
 * ⁠      1
 * ⁠    /   \
 * ⁠   2     3
 * ⁠    \  
 * ⁠     4 
 * 
 * 输出: "1(2()(4))(3)"
 * 
 * 解释: 和第一个示例相似，
 * 除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。
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
//dfs
// class Solution {

//     public String tree2str(TreeNode t) {
//         if (t == null) return "";

//         if (t.left == null && t.right == null) return String.valueOf(t.val);

//         if (t.right == null) {
//             return t.val + "(" + tree2str(t.left) + ")";
//         }
//         return t.val + "(" + tree2str(t.left) + ")(" + tree2str(t.right) + ")";
//     }
// }
// iterative
// class Solution {

//     public String tree2str(TreeNode t) {
//         if (t == null) return "";
//         Deque<TreeNode> stack = new LinkedList<>();
//         stack.push(t);
//         Set<TreeNode> visited = new HashSet<>();
//         StringBuilder ans = new StringBuilder(8);

//         while (!stack.isEmpty()) {
//             t = stack.peek();
//             if (visited.contains(t)) {
//                 stack.pop();
//                 ans.append(")");
//                 continue;
//             } 
//             visited.add(t);
//             ans.append("(" + t.val);
//             if (t.left == null && t.right != null) {
//                 ans.append("()");
//             }
//             if (t.right != null) {
//                 stack.push(t.right);
//             }
//             if (t.left != null) {
//                 stack.push(t.left);
//             }
//         }
//         return ans.substring(1, ans.length() - 1);
//     }
// }
class Solution {
    private void dfs(TreeNode root, TreeNode pre, StringBuilder ans) {
        if (root == null ) return;

        // if left child is empty add ()
        if (pre != null && pre.left == null && pre.right == root) {
            ans.append("()");
        }
        ans.append("(").append(root.val);
        pre = root;
        dfs(root.left, pre, ans);
        dfs(root.right, pre, ans);
        ans.append(")");
        
    }
    public String tree2str(TreeNode t) {
        if (t == null) return "";
        StringBuilder ans = new StringBuilder();
        dfs(t, null, ans);
        return ans.substring(1, ans.length() - 1);
    }
}
// @lc code=end

