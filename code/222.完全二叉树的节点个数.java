import javax.swing.tree.TreeNode;

/*
 * @lc app=leetcode.cn id=222 lang=java
 *
 * [222] 完全二叉树的节点个数
 *
 * https://leetcode-cn.com/problems/count-complete-tree-nodes/description/
 *
 * algorithms
 * Medium (71.84%)
 * Likes:    252
 * Dislikes: 0
 * Total Accepted:    35.1K
 * Total Submissions: 48.3K
 * Testcase Example:  '[1,2,3,4,5,6]'
 *
 * 给出一个完全二叉树，求出该树的节点个数。
 * 
 * 说明：
 * 
 * 
 * 完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第
 * h 层，则该层包含 1~ 2^h 个节点。
 * 
 * 示例:
 * 
 * 输入: 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   3
 * ⁠/ \  /
 * 4  5 6
 * 
 * 输出: 6
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

 // recursive version
// class Solution {
//     private int getHeight(TreeNode root) {
//         int h = 0;
//         while (root != null) {
//             root = root.left;
//             h++;
//         }
//         return h;
//     }
//     public int countNodes(TreeNode root) {
//         if (root == null) return 0;
//         int leftHeight = getHeight(root.left);
//         int rightHeight = getHeight(root.right);

//         // left is complete tree
//         if (leftHeight == rightHeight) {
//             return countNodes(root.right) + (1 << leftHeight);
//         } else {
//             // left > right, right is complete tree
//             return countNodes(root.left) + (1 << rightHeight);
//         }
//     }
// }

class Solution {
    
    private boolean exist(int target, int h, TreeNode root) {
        int l = 0, r = (1 << h) - 1;
        while (h-- > 0) {
            int mid = l + (r - l) / 2;
            if (target <= mid) {
                root = root.left;
                r = mid;    
            } else {
                root = root.right;
                l = mid + 1;
            }
        }

        return root != null;
    }

    private int getHeight(TreeNode root) {
        int h = 0;
        while (root.left != null) {
            root = root.left;
            h++;
        }
        return h;
    }

    public int countNodes(TreeNode root) {
        if (root == null) return 0;
        // we know that the nodes will be (2^n - 1) + leaves
        int h = getHeight(root);
        int l = 1, r = (1 << h) - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;

            if (exist(mid, h, root)) {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

        return (1 << h) - 1 + l;
    }
}
// @lc code=end

