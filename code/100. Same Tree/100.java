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
    public boolean isSameTree(TreeNode p, TreeNode q) {
        if (p == null && q == null) return true;
        if ((p != null && q == null) || (p == null && q != null)) return false;

        
        return p.val == q.val && isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}



// bfs
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
    private static boolean isEqual(TreeNode p, TreeNode q) {

        if (p == null && q == null) return true;
        if ((p != null && q == null) || (p == null && q != null)) return false;
        if (p.val != q.val) return false;
        return true;
    }
    
    public boolean isSameTree(TreeNode p, TreeNode q) {
        LinkedList<TreeNode> pqueue = new LinkedList<>();
        LinkedList<TreeNode> qqueue = new LinkedList<>();
        
        pqueue.add(p);
        qqueue.add(q);
        
        while (!pqueue.isEmpty()) {
            TreeNode t1 = pqueue.pop();
            TreeNode t2 = qqueue.pop();
            
            if (!isEqual(t1, t2)) return false;
            
            if (t1 != null) {
                pqueue.add(t1.left);
                pqueue.add(t1.right);
            }
            if (t2 != null) {

                qqueue.add(t2.left);
                qqueue.add(t2.right);
            } 
        
        }
        return true;
    }
}