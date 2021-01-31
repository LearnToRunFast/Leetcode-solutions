import java.util.ArrayList;
import java.util.Deque;
import java.util.HashSet;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=140 lang=java
 *
 * [140] 单词拆分 II
 *
 * https://leetcode-cn.com/problems/word-break-ii/description/
 *
 * algorithms
 * Hard (44.30%)
 * Likes:    406
 * Dislikes: 0
 * Total Accepted:    40K
 * Total Submissions: 90.2K
 * Testcase Example:  '"catsanddog"\n["cat","cats","and","sand","dog"]'
 *
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典
 * wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
 * 
 * 说明：
 * 
 * 
 * 分隔时可以重复使用字典中的单词。
 * 你可以假设字典中没有重复的单词。
 * 
 * 
 * 示例 1：
 * 
 * 输入:
 * s = "catsanddog"
 * wordDict = ["cat", "cats", "and", "sand", "dog"]
 * 输出:
 * [
 * "cats and dog",
 * "cat sand dog"
 * ]
 * 
 * 
 * 示例 2：
 * 
 * 输入:
 * s = "pineapplepenapple"
 * wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
 * 输出:
 * [
 * "pine apple pen apple",
 * "pineapple pen apple",
 * "pine applepen apple"
 * ]
 * 解释: 注意你可以重复使用字典中的单词。
 * 
 * 
 * 示例 3：
 * 
 * 输入:
 * s = "catsandog"
 * wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出:
 * []
 * 
 * 
 */

// @lc code=start
// class Solution {
//     public List<String> wordBreak(String s, List<String> wordDict) {
//         Set<String> dict = new HashSet<>(wordDict);

//         boolean[] dp = new boolean[s.length() + 1];
//         dp[0] = true;
//         for (int i = 1; i <= s.length(); i++) {
//             for (int j = 0; j < i; j++) {
//                 if (dp[j] && dict.contains(s.substring(j, i))) {
//                     dp[i] = true;
//                     break; // we can at least one true on i
//                 }
//             }
//         }
//         if (!dp[s.length()]) return new ArrayList<>();

//         List<List<String>> strs = new ArrayList<>(s.length() + 1);

//         for (int i = 0; i <= s.length(); i++) {
//             strs.add(new ArrayList<>());
//         }

//         for (int i = 1; i <= s.length(); i++) {
//             if (!dp[i]) continue;

//             for (int j = 0; j < i; j++) {
//                 if (!dp[j]) continue;
//                 String word = s.substring(j, i);
//                 if (!dict.contains(word)) continue;

//                 if (strs.get(j).size() == 0) {
//                     strs.get(i).add(word);
//                 } else {
//                     for (String str : strs.get(j)) {
//                         strs.get(i).add(new StringBuilder(str).append(" ").append(word).toString());
//                     }
//                 }

//             }
//         }
//         return strs.get(s.length());
//     }
// }
class Solution {
    public List<String> wordBreak(String s, List<String> wordDict) {
        Set<String> dict = new HashSet<>(wordDict);

        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true;
        for (int r = 1; r <= s.length(); r++) {
            for (int l = 0; l < r; l++) {
                if (dp[l] && dict.contains(s.substring(l, r))) {
                    dp[r] = true;
                    break; // we can at least one true on i
                }
            }
        }
        List<String> ans = new ArrayList<>();
        if (!dp[s.length()]) return ans;
        Deque<String> words = new LinkedList<>();
        dfs(s, s.length(), dict, dp, words, ans);
        return ans;
    }
    private void dfs(String s, int r, Set<String> dict, boolean[] dp, Deque<String>words, List<String> ans) {
        if (r == 0) {
            ans.add(String.join(" ", words));
        }

        for (int l = r - 1; l >= 0; l--) {
            String suffix = s.substring(l, r);
            if (dp[l] && dict.contains(suffix)) {
                words.addFirst(suffix);
                dfs(s, l, dict, dp, words, ans);
                words.pollFirst();
            }
        }
    }
}
// @lc code=end

