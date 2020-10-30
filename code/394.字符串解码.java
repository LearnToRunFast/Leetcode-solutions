import java.util.Stack;

/*
 * @lc app=leetcode.cn id=394 lang=java
 *
 * [394] 字符串解码
 *
 * https://leetcode-cn.com/problems/decode-string/description/
 *
 * algorithms
 * Medium (53.06%)
 * Likes:    550
 * Dislikes: 0
 * Total Accepted:    66.7K
 * Total Submissions: 125.1K
 * Testcase Example:  '"3[a]2[bc]"'
 *
 * 给定一个经过编码的字符串，返回它解码后的字符串。
 * 
 * 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
 * 
 * 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
 * 
 * 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "3[a]2[bc]"
 * 输出："aaabcbc"
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "3[a2[c]]"
 * 输出："accaccacc"
 * 
 * 
 * 示例 3：
 * 
 * 输入：s = "2[abc]3[cd]ef"
 * 输出："abcabccdcdcdef"
 * 
 * 
 * 示例 4：
 * 
 * 输入：s = "abc3[cd]xyz"
 * 输出："abccdcdcdxyz"
 * 
 * 
 */

// @lc code=start
// recursion
// class Solution {

//     private String string;
//     private int ptr;
//     private int len;

//     private boolean isNumber(char c) {
//         int res = c - '0';
//         return res >= 0 && res <= 9;
//     }

//     private int findOpenOrClose(String s, int i, char c) {
//         while (i < s.length() && s.charAt(i) != c) {
//             i++;
//         }
//         return i;
//     }

//     private StringBuilder decodeString() {
//         StringBuilder sb = new StringBuilder();
//         if (ptr >= len || string.charAt(ptr) == ']') return sb;

//         char curr = string.charAt(ptr);
 

//         if (isNumber(curr)) {
//             int idx = findOpenOrClose(string, ptr + 1, '[');
//             int repTime = Integer.parseInt(string.substring(ptr, idx));
//             ptr = idx + 1;
//             StringBuilder tmp = decodeString();
//             ptr++;
//             while (repTime-- > 0) {
//                 sb.append(tmp);
//             }
            
//         } else {
//             sb.append(curr);
//             ptr++;
//         }
//         return sb.append(decodeString());
//     }
//     public String decodeString(String s) {
//         string = s;
//         ptr = 0;
//         len = s.length();
//         return decodeString().toString();

//     }
// }
class Solution {

    private boolean isNumber(char c) {
        int res = c - '0';
        return res >= 0 && res <= 9;
    }

    public String decodeString(String s) {
        Stack<String> stackAns = new Stack<>();
        Stack<Integer> stackNum = new Stack<>();

        StringBuilder ans = new StringBuilder();
        int num = 0;

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '[') {
                stackAns.push(ans.toString());
                stackNum.push(num);
                ans = new StringBuilder();
                num = 0;
            } else if (c == ']') {
                StringBuilder tmp = new StringBuilder();
                tmp.append(stackAns.pop());
                int count = stackNum.pop();
                while (count-- > 0) {
                    tmp.append(ans);                    
                }
                ans = tmp;
            } else if (isNumber(c)) {
                num = num * 10  + Integer.parseInt(String.valueOf(c));
            } else {
                ans.append(c);
            } 
        }
        return ans.toString();
    }
}
// @lc code=end

