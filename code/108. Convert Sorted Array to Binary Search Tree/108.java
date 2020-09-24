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
    public TreeNode sortedArrayToBST(int[] nums) {
        return sortedArrayToBST(nums, 0, nums.length);
    }
    private static TreeNode sortedArrayToBST(int[] nums, int start, int end) {
        if (start == end) return null;
        
        int mid = (start + end) >>> 1;
        TreeNode t = new TreeNode(nums[mid]);
        t.left = sortedArrayToBST(nums, start, mid);
        t.right = sortedArrayToBST(nums, mid + 1, end);
        return t;
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
    public TreeNode sortedArrayToBST(int[] nums) {
        return sortedArrayToBST(nums, 0, nums.length - 1);
    }
    private static TreeNode sortedArrayToBST(int[] nums, int start, int end) {
        if (start > end) return null;
        
        int mid = (start + end) >>> 1;
        TreeNode t = new TreeNode(nums[mid]);
        t.left = sortedArrayToBST(nums, start, mid - 1);
        t.right = sortedArrayToBST(nums, mid + 1, end);
        return t;
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
    public TreeNode sortedArrayToBST(int[] nums) {
        if (nums == null || nums.length == 0) return null;
        
        LinkedList<TreeNode> queue = new LinkedList<>();
        int mid = nums.length >>> 1;
        queue.add(nums[mid]);
        while (!queue.isEmpty()) {
            
        }
    }

}