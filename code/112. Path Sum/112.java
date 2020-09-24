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

    public boolean hasPathSum(TreeNode root, int sum) {
        if (root == null) return false;

        if (root.left == null && root.right == null) return (sum - root.val) == 0;
        boolean left = hasPathSum(root.left, sum - root.val);
        boolean right = hasPathSum(root.right, sum - root.val);
        return left || right ;
        
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

    public boolean hasPathSum(TreeNode root, int sum) {
        if (root == null) return false;
        LinkedList<TreeNode> queue = new LinkedList<>();
        LinkedList<Integer> currSum = new LinkedList<>();
        queue.add(root);
        currSum.add(sum);

        while (!queue.isEmpty()) {
            TreeNode t = queue.pop();
            int temp = currSum.pop() - t.val;

            if (t.left == null && t.right == null && temp == 0) return true;

            if (t.left != null) {
                queue.add(t.left);
                currSum.add(temp);
            }

            if (t.right != null) {
                queue.add(t.right);
                currSum.add(temp);
            }
        }
        return false;
        
    }
}