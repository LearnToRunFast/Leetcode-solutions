/*
 * @lc app=leetcode.cn id=29 lang=java
 *
 * [29] 两数相除
 *
 * https://leetcode-cn.com/problems/divide-two-integers/description/
 *
 * algorithms
 * Medium (20.10%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    64.7K
 * Total Submissions: 321.1K
 * Testcase Example:  '10\n3'
 *
 * 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
 * 
 * 返回被除数 dividend 除以除数 divisor 得到的商。
 * 
 * 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
 * -2
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: dividend = 10, divisor = 3
 * 输出: 3
 * 解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
 * 
 * 示例 2:
 * 
 * 输入: dividend = 7, divisor = -3
 * 输出: -2
 * 解释: 7/-3 = truncate(-2.33333..) = -2
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 被除数和除数均为 32 位有符号整数。
 * 除数不为 0。
 * 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。
 * 
 * 
 */

// @lc code=start
class Solution {
    public int divide(int dividend, int divisor) {
        // edge case
        if (dividend == Integer.MIN_VALUE && divisor == -1) return Integer.MAX_VALUE;

        // check left most bits
        //somehow >> 1 will causing the left most bit become 1 instead of 0
        boolean sameSign = (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0);
        int ans = 0;
        dividend = dividend > 0 ? ~(dividend - 1) : dividend;
        divisor = divisor > 0 ? ~(divisor - 1) : divisor;

        while(dividend <= divisor) {
            int tmp = divisor;
            int factor = 1;

            // edge case if divisor = -1
            // dividend - tmp <= tmp will be totally different from dividend <= tmp + tmp
            // here as dividend and divisor are both negative
            // so dividend - tmp will never overflow
            while(dividend - tmp <= tmp) {
                tmp <<= 1;
                factor <<= 1;
            }
            
            dividend -= tmp;
            ans += factor;
        }

        return sameSign ? ans : -ans;

    }
}


// @lc code=end

