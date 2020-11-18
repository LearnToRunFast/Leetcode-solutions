import java.util.LinkedList;

/*
 * @lc app=leetcode.cn id=402 lang=java
 *
 * [402] 移掉K位数字
 *
 * https://leetcode-cn.com/problems/remove-k-digits/description/
 *
 * algorithms
 * Medium (32.31%)
 * Likes:    449
 * Dislikes: 0
 * Total Accepted:    49.2K
 * Total Submissions: 150.5K
 * Testcase Example:  '"1432219"\n3'
 *
 * 给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
 * 
 * 注意:
 * 
 * 
 * num 的长度小于 10002 且 ≥ k。
 * num 不会包含任何前导零。
 * 
 * 
 * 示例 1 :
 * 
 * 
 * 输入: num = "1432219", k = 3
 * 输出: "1219"
 * 解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
 * 
 * 
 * 示例 2 :
 * 
 * 
 * 输入: num = "10200", k = 1
 * 输出: "200"
 * 解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
 * 
 * 
 * 示例 3 :
 * 
 * 
 * 输入: num = "10", k = 2
 * 输出: "0"
 * 解释: 从原数字移除所有的数字，剩余为空就是0。
 * 
 * 
 */

// @lc code=start
class Solution {
    public String removeKdigits(String num, int k) {
        int n = num.length();
        if (k >= n) {
            return "0";
        }
        StringBuilder sb = new StringBuilder();

        LinkedList<Character> stack = new LinkedList<>();
        int remain = n - k;
        for (int i = 0; i < n; i++) {
            
            while (k > 0 && !stack.isEmpty() && stack.peek() > num.charAt(i)) {
                stack.pop();
                k--;
            }

            stack.push(num.charAt(i));
        }

        while (!stack.isEmpty() && stack.peekLast() == '0') {
            stack.pollLast();
            remain--;
        }

        int curr = 0;
        while (!stack.isEmpty() && curr < remain) {
            sb.append(stack.pollLast());
            curr++;
        }

        return sb.length() > 0 ? sb.toString() : "0";
    }
}
// @lc code=end

