import java.util.HashMap;
import java.util.Map;

/*
 * @lc app=leetcode.cn id=87 lang=java
 *
 * [87] 扰乱字符串
 *
 * https://leetcode-cn.com/problems/scramble-string/description/
 *
 * algorithms
 * Hard (48.64%)
 * Likes:    176
 * Dislikes: 0
 * Total Accepted:    16.4K
 * Total Submissions: 33.6K
 * Testcase Example:  '"great"\n"rgeat"'
 *
 * 给定一个字符串 s1，我们可以把它递归地分割成两个非空子字符串，从而将其表示为二叉树。
 * 
 * 下图是字符串 s1 = "great" 的一种可能的表示形式。
 * 
 * ⁠   great
 * ⁠  /    \
 * ⁠ gr    eat
 * ⁠/ \    /  \
 * g   r  e   at
 * ⁠          / \
 * ⁠         a   t
 * 
 * 
 * 在扰乱这个字符串的过程中，我们可以挑选任何一个非叶节点，然后交换它的两个子节点。
 * 
 * 例如，如果我们挑选非叶节点 "gr" ，交换它的两个子节点，将会产生扰乱字符串 "rgeat" 。
 * 
 * ⁠   rgeat
 * ⁠  /    \
 * ⁠ rg    eat
 * ⁠/ \    /  \
 * r   g  e   at
 * ⁠          / \
 * ⁠         a   t
 * 
 * 
 * 我们将 "rgeat” 称作 "great" 的一个扰乱字符串。
 * 
 * 同样地，如果我们继续交换节点 "eat" 和 "at" 的子节点，将会产生另一个新的扰乱字符串 "rgtae" 。
 * 
 * ⁠   rgtae
 * ⁠  /    \
 * ⁠ rg    tae
 * ⁠/ \    /  \
 * r   g  ta  e
 * ⁠      / \
 * ⁠     t   a
 * 
 * 
 * 我们将 "rgtae” 称作 "great" 的一个扰乱字符串。
 * 
 * 给出两个长度相等的字符串 s1 和 s2，判断 s2 是否是 s1 的扰乱字符串。
 * 
 * 示例 1:
 * 
 * 输入: s1 = "great", s2 = "rgeat"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: s1 = "abcde", s2 = "caebd"
 * 输出: false
 * 
 */

// @lc code=start
// class Solution {
//     private boolean helper(String s1, String s2, Map<String, Integer> map) {
//         int preResult = map.getOrDefault(s1 + "#" + s2, -1);

//         if (preResult != -1) return preResult == 1 ? true : false;


//         if (s1.length() != s2.length()) {
//             map.put(s1 + "#" + s2, 0);
//             return false;
//         }
//         if (s1.equals(s2)) {
//             map.put(s1 + "#" + s2, 1);
//             return true;            
//         }
//         int[] letters = new int[26];
//         for (int i = 0; i < s1.length(); i++) {
//             letters[s1.charAt(i) - 'a']++;
//             letters[s2.charAt(i) - 'a']--;
//         }
//         for (int i = 0; i < 26;  i++) {
//             if (letters[i] != 0) {
//                 map.put(s1 + "#" + s2, 0);
//                 return false;
//             }
//         }
//         for (int i = 1; i < s1.length(); i++) {
//             if (isScramble(s1.substring(0, i), s2.substring(0, i)) && isScramble(s1.substring(i), s2.substring(i))) {
//                 map.put(s1 + "#" + s2, 1);
//                 return true;
//             }
//             // first ith of s1 compare to tail ith of s2
//             if (isScramble(s1.substring(0, i), s2.substring(s2.length() - i)) &&
//                 isScramble(s1.substring(i), s2.substring(0, s2.length() - i))) {
//                 map.put(s1 + "#" + s2, 1);
//                 return true;
//             }
//         }
//         map.put(s1 + "#" + s2, 0);
//         return false;
//     }
//     public boolean isScramble(String s1, String s2) {
//         Map<String, Integer> map = new HashMap<>();
//         return helper(s1, s2, map);
//     }
// }
class Solution {
    public boolean isScramble(String s1, String s2) {
        if (s1.length() != s2.length()) {
            return false;
        }
        if (s1.equals(s2)) {
            return true;            
        }
        int[] letters = new int[26];
        for (int i = 0; i < s1.length(); i++) {
            letters[s1.charAt(i) - 'a']++;
            letters[s2.charAt(i) - 'a']--;
        }
        for (int i = 0; i < 26;  i++) {
            if (letters[i] != 0) {
                return false;
            }
        }
        for (int i = 1; i < s1.length(); i++) {
            if (isScramble(s1.substring(0, i), s2.substring(0, i)) && isScramble(s1.substring(i), s2.substring(i))) {
                return true;
            }
            // first ith of s1 compare to tail ith of s2
            if (isScramble(s1.substring(0, i), s2.substring(s2.length() - i)) &&
                isScramble(s1.substring(i), s2.substring(0, s2.length() - i))) {
                return true;
            }
        }
        return false;
    }
}
// @lc code=end

