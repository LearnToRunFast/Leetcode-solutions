/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    
    private static int recursion(TreeNode t1) {
        if (t1 == null) return 0;
        int left = recursion(t1.left);
        if (left == -1) return -1;
        int right = recursion(t1.right);
        if (right == -1) return -1;
        return Math.abs(left - right) > 1 ? -1 : Math.max(left, right) + 1;
    } 
    public boolean isBalanced(TreeNode root) {
        
        int ans = recursion(root);
        return ans != -1 ? true : false;
    }
}
