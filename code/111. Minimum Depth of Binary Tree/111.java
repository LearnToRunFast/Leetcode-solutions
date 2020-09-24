
class Solution {
    public int minDepth(TreeNode root) {
        if (root == null) return 0;
        // if it's both empty, then its a leaf, + 1;
        if (root.left == null && root.right == null) return 1;
        
        int left = minDepth(root.left);
        int right = minDepth(root.right);
        
        //  root.left == null || root.right == null means one of is empty, so left + right will be 0 + sth
        return root.left == null || root.right == null ? left + right + 1: Math.min(left, right) + 1;
        
    }
}

import javafx.util.Pair;
class Solution {
  public int minDepth(TreeNode root) {
    if (root == null) return 0;
      
    LinkedList<Pair<TreeNode, Integer>> queue = new LinkedList<>();


    queue.add(new Pair(root, 1));

    int current_depth = 0;
    while (!queue.isEmpty()) {
      Pair<TreeNode, Integer> current = queue.poll();
      root = current.getKey();
      current_depth = current.getValue();
      // if current is leap
      if ((root.left == null) && (root.right == null)) break;
        
      if (root.left != null) {
        queue.add(new Pair(root.left, current_depth + 1));
      }
      if (root.right != null) {
        queue.add(new Pair(root.right, current_depth + 1));
      }
    }
    return current_depth;
  }
}

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { 
       val = x; }
 * }
 */
class Solution {
    public int minDepth(TreeNode root) {
        if (root == null) return 0;
        LinkedList<TreeNode> queue = new LinkedList<>();
        
        queue.add(root);
        int height = 1;
        
        while(!queue.isEmpty()) {
            
            int len = queue.size();
            
            for (int i = 0; i < len; i++) {
                TreeNode t = queue.pop();
                // this is a leaf
                if (t.left == null && t.right == null) return height;
                // if its not a leaf, going to search down
                if (t.left != null) queue.add(t.left);
                if (t.right != null) queue.add(t.right);
            }
            height++;
        }
        return height;
    }
}
