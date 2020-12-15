import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=501 lang=java
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (49.69%)
 * Likes:    250
 * Dislikes: 0
 * Total Accepted:    44.4K
 * Total Submissions: 89.3K
 * Testcase Example:  '[1,null,2,2]'
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 * 
 * 假定 BST 有如下定义：
 * 
 * 
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 * 
 * 
 * 例如：
 * 给定 BST [1,null,2,2],
 * 
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 * 
 * 
 * 返回[2].
 * 
 * 提示：如果众数超过1个，不需考虑输出顺序
 * 
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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
// morris
// TreeNode predecessor = null;
// while (root != null) {
//     if (root.left == null) {
//         // action here
//         root = root.right;
//         continue;
//     }
//     // if root.left != null
//     // find predecessor and it stops correctly if we visited
//     predecessor = root.left;
//     while (predecessor.right != null && predecessor.right != root) {
//         predecessor = predecessor.right;
//     }
//     // we have two scenarios
//     if (predecessor.right == null) {
//         predecessor.right = root;
//         root = root.left;
//     } else {
//         predecessor.right = null;
//         // action here
//         // already access all the nodes on left
//         root = root.right;
//     }
// }
// class Solution {
//     private int base, count, maxCount;
//     List<Integer> ans = new ArrayList<>();

//     private void update(int num) {
//         if (num == base) {
//             count++;
//         } else {
//             count = 1;
//             base = num;
//         }
//         if (count == maxCount) {
//             ans.add(num);
//         } else if (count > maxCount) {
//             maxCount = count;
//             ans.clear();
//             ans.add(num);
//         }
//     }

//     public int[] findMode(TreeNode root) {
//         TreeNode predecessor = null;
//         while (root != null) {
//             if (root.left == null) {
//                 // action here
//                 update(root.val);
//                 root = root.right;
//                 continue;
//             }
//             // if root.left != null
//             // find predecessor and it stops correctly if we visited
//             // predecessor before
//             predecessor = root.left;
//             while (predecessor.right != null && predecessor.right != root) {
//                 predecessor = predecessor.right;
//             }
//             // we have two scenarios
//             if (predecessor.right == null) {
//                 predecessor.right = root;
//                 root = root.left;
//             } else {
//                 predecessor.right = null;
//                 // action here
//                 update(root.val);
//                 // already access all the nodes on left
//                 root = root.right;
//             }
//         }
//         // convert list to array
//         int[] mode = new int[ans.size()];
//         for (int i = 0; i < ans.size(); i++) {
//             mode[i] = ans.get(i);
//         }
//         return mode;
//     }
// }
class Solution {
    private int base, count, maxCount;
    List<Integer> ans = new ArrayList<>();

    private void update(int num) {
        if (num == base) {
            count++;
        } else {
            count = 1;
            base = num;
        }
        if (count == maxCount) {
            ans.add(num);
        } else if (count > maxCount) {
            maxCount = count;
            ans.clear();
            ans.add(num);
        }
    }
    private void dfs(TreeNode root) {
        if (root == null) return;
        dfs(root.left);
        update(root.val);
        dfs(root.right);
    }
    public int[] findMode(TreeNode root) {
        dfs(root);
        // convert list to array
        int[] mode = new int[ans.size()];
        for (int i = 0; i < ans.size(); i++) {
            mode[i] = ans.get(i);
        }
        return mode;
    }
}
// @lc code=end

