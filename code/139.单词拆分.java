import java.util.ArrayList;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=139 lang=java
 *
 * [139] 单词拆分
 *
 * https://leetcode-cn.com/problems/word-break/description/
 *
 * algorithms
 * Medium (47.45%)
 * Likes:    744
 * Dislikes: 0
 * Total Accepted:    101.6K
 * Total Submissions: 209.9K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 * 
 * 说明：
 * 
 * 
 * 拆分时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 * 
 * 
 * 示例 1：
 * 
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
 * 
 * 
 * 示例 2：
 * 
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
 * 注意你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 * 
 * 
 */

// @lc code=start
public class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> set = new HashSet(wordDict);
        int n = s.length();

        List<Integer> dp = new ArrayList<>();
        dp.add(0);
        for (int end = 1; end <= n; end++) {
            for (int start : dp) {
                if (set.contains(s.substring(start, end))) {
                    dp.add(end);
                    break;
                }
            }
        }
        return dp.get(dp.size() - 1) == n;
    }
}

// @lc code=end

