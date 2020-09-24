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

    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;
        
        LinkedList<TreeNode> queue = new LinkedList<>();
        
        queue.add(root.left);
        queue.add(root.right);
        
        while (!queue.isEmpty()) {
            
            TreeNode t1 = queue.pop();
            TreeNode t2 = queue.pop();
            
            if (t1 == null && t2 == null) continue;
            if (t1 == null || t2 == null) return false;
            if (t1.val != t2.val) return false;
            
            queue.add(t1.left);
            queue.add(t2.right);
            queue.add(t1.right);
            queue.add(t2.left);
            
        }
        return true;
    }
}


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
    
    private static boolean isSame(TreeNode t1, TreeNode t2) {
        if (t1 == null && t2 == null) return true; 
        if (t1 == null || t2 == null) return false;
        // why not t1.val == t2.val return true? because if they are equal, it will continue to sub tree.
        return t1.val == t2.val && isSame(t1.left, t2.right) && isSame(t1.right, t2.left);
        
    }
    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;
        return isSame(root.left, root.right);
    }
}