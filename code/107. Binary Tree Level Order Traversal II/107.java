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
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        LinkedList<TreeNode> queue = new LinkedList<>();
        
        queue.add(root);
        List<List<Integer>> ans = new ArrayList<>();
        while (!queue.isEmpty()) {
            ArrayList<Integer> arraylist = new ArrayList<>();
            
            int len = queue.size();
            for (int i = 0; i < len; i++) {
                TreeNode t = queue.pop();
                if (t == null) continue;
                arraylist.add(t.val);
                queue.add(t.left);
                queue.add(t.right);
            }
            
            if (arraylist.size() > 0) ans.add(arraylist);
            
        }
        
        Collections.reverse(ans);
        return ans;
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
    private static void DFS(TreeNode t, int level, List<List<Integer>> ans) {
        if (t == null) return;
        
        if (ans.size() <= level) ans.add(0, new ArrayList<Integer>());
        
        ans.get(ans.size() - 1 - level).add(t.val);
        
        DFS(t.left, level + 1, ans);
        DFS(t.right, level + 1, ans);
        
    }
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> ans = new ArrayList<>();
        
        DFS(root, 0, ans);
        
        return ans;
    }
}