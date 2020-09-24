/*
 * @lc app=leetcode.cn id=125 lang=java
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (41.07%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    56.5K
 * Total Submissions: 137.2K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 * 
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 * 
 * 示例 1:
 * 
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 * 
 * 
 * 示例 2:
 * 
 * 输入: "race a car"
 * 输出: false
 * 
 * 
 */

// @lc code=start
class Solution {
    private static boolean isAlnum(char c) {
        if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
            return true;
        }
        return false;
    }

    public boolean isPalindrome(String s) {
        int i = 0, j = s.length() - 1;
        while (i < j) {
            while (i < j && !isAlnum(s.charAt(i)))
                i++;
            while (i < j && !isAlnum(s.charAt(j)))
                j--;
            if (Character.toLowerCase(s.charAt(i++)) != Character.toLowerCase(s.charAt(j--)))
                return false;

        }
        return true;
    }
}
// @lc code=end
