import java.util.Deque;
import java.util.LinkedList;
import java.util.List;

import javax.swing.tree.TreeNode;

import sun.awt.geom.AreaOp.IntOp;

/*
 * @lc app=leetcode.cn id=449 lang=java
 *
 * [449] 序列化和反序列化二叉搜索树
 *
 * https://leetcode-cn.com/problems/serialize-and-deserialize-bst/description/
 *
 * algorithms
 * Medium (53.99%)
 * Likes:    135
 * Dislikes: 0
 * Total Accepted:    9.3K
 * Total Submissions: 17.3K
 * Testcase Example:  '[2,1,3]'
 *
 * 序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
 * 
 * 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。
 * 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
 * 
 * 编码的字符串应尽可能紧凑。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [2,1,3]
 * 输出：[2,1,3]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = []
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数范围是 [0, 10^4]
 * 0 
 * 题目数据 保证 输入的树是一棵二叉搜索树。
 * 
 * 
 * 
 * 
 * 注意：不要使用类成员/全局/静态变量来存储状态。 你的序列化和反序列化算法应该是无状态的。
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
// base on BST properties
// public class Codec {
//     private void serialize(TreeNode root, StringBuilder sb) {
//         if (root == null) {
//             return;
//         }
//         sb.append(root.val).append(",");
//         serialize(root.left, sb);
//         serialize(root.right, sb);
//     }
//     // Encodes a tree to a single string.
//     public String serialize(TreeNode root) {
//         StringBuilder sb = new StringBuilder();
//         serialize(root, sb);
//         return sb.toString();
//     }
    
//     private TreeNode deserialize(String[] token, int lo, int hi) {
//         if (lo > hi) return null;
//         TreeNode root = new TreeNode(Integer.valueOf(token[lo]));

//         int mid = hi + 1;
//         for (int i = lo + 1; i <= hi; i++) {
//             if (Integer.valueOf(token[i]) > root.val) {
//                 mid = i;
//                 break;
//             }
//         }
//         root.left = deserialize(token, lo + 1, mid - 1);
//         root.right = deserialize(token, mid, hi);
//         return root;
//     }
//     // Decodes your encoded data to tree.
//     public TreeNode deserialize(String data) {
//         if (data == null || data.length() == 0) return null;
//         String[] token = data.split(",");
//         return deserialize(token, 0, token.length - 1);

//     }
// }
// with null in serializer
// public class Codec {
//     private void serialize(TreeNode root, StringBuilder sb) {
//         if (root == null) {
//             sb.append("#").append(",");
//             return;
//         }
//         sb.append(root.val).append(",");
//         serialize(root.left, sb);
//         serialize(root.right, sb);
//     }
//     // Encodes a tree to a single string.
//     public String serialize(TreeNode root) {
//         StringBuilder sb = new StringBuilder();
//         serialize(root, sb);
//         return sb.toString();
//     }
    
//     private int step;
//     private TreeNode deserialize(String[] token) {
//         if (step == token.length) return null;
//         TreeNode root = null;
//         if (!token[step].equals("#")) {
//             root = new TreeNode(Integer.valueOf(token[step]));
//             step++;
//             root.left = deserialize(token);
//             step++;
//             root.right = deserialize(token);
//         }
//         return root;

//     }
//     // Decodes your encoded data to tree.
//     public TreeNode deserialize(String data) {
//         if (data.charAt(0) == '#') return null;
//         String[] token = data.split(",");
//         return deserialize(token);

//     }
// }
// without null in serializer
public class Codec {
    // post order , why post order ?  during deserialize, remove at tail is more fast
    // than remove at head for list
    private void serialize(TreeNode root, StringBuilder sb) {
        if (root == null) {
            return;
        }
        serialize(root.left, sb);
        serialize(root.right, sb);
        sb.append(root.val).append(",");
    }
    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        serialize(root, sb);
        return sb.toString();
    }
    
    private TreeNode deserialize(Deque<Integer> tokens, int min, int max) {
        if (tokens.isEmpty()) return null;
        int value = tokens.peekLast();
        if (min > value || max < value) return null;
        TreeNode root = new TreeNode(tokens.pollLast());
        root.right = deserialize(tokens, value, max);
        root.left = deserialize(tokens, min, value);

        return root;

    }
    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if (data.length() == 0) return null;
        String[] tokens = data.split(",");
        Deque<Integer> intTokens = new LinkedList<>();
        for (String token : tokens) {
            intTokens.add(Integer.valueOf(token));
        }

        return deserialize(intTokens, Integer.MIN_VALUE, Integer.MAX_VALUE);

    }
}
// Your Codec object will be instantiated and called as such:
// Codec ser = new Codec();
// Codec deser = new Codec();
// String tree = ser.serialize(root);
// TreeNode ans = deser.deserialize(tree);
// return ans;
// @lc code=end

