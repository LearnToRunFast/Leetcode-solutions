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
// public class Solution {
//     public boolean wordBreak(String s, List<String> wordDict) {
//         Set<String> set = new HashSet(wordDict);
//         int n = s.length();

//         List<Integer> dp = new ArrayList<>();
//         dp.add(0);
//         for (int end = 1; end <= n; end++) {
//             for (int start : dp) {
//                 if (set.contains(s.substring(start, end))) {
//                     dp.add(end);
//                     break;
//                 }
//             }
//         }
//         return dp.get(dp.size() - 1) == n;
//     }
// }

public class Solution {
    class TrieNode {
        private final int R = 26;
        private boolean isEnd;
        private TrieNode[] links;

        public TrieNode() {
            links = new TrieNode[R];
        }

        public boolean containsKey(char c) {
            return links[c - 'a'] != null;
        }

        public TrieNode get(char c) {
            return links[c - 'a'];
        }
        
        public void put(char c) {
            links[c - 'a'] = new TrieNode();
        }

        public void setEnd() {
            this.isEnd = true;
        }

        public boolean match() {
            return this.isEnd;
        }

    }
    class Trie {
        private TrieNode root;
        private boolean searchResult;
        private int[] isVisited; 
        private String word;
        private int n;

        public Trie() {
            root = new TrieNode();
        }

        public void insert(String word) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (!node.containsKey(c)) {
                    node.put(c);
                }
                node = node.get(c);
            }
            node.setEnd();
        }

        private void backtrack(int idx, TrieNode curr) {
            if (searchResult || idx == n || isVisited[idx] != 0) return;

            isVisited[idx] = 1;
            for (int i = idx; i < n && !searchResult; i++) {
                char c = word.charAt(i);

                if (!curr.containsKey(c)) {
                    return;
                }
                curr = curr.get(c);
                if (curr.match()) {
                    if (i == n - 1) {
                        searchResult = true;
                        return;
                    }

                    if(i < n - 1 && isVisited[i + 1] == 0) {
                        backtrack(i + 1, root);
                    }
                }
            }
        }
        public boolean search(String s) {
            word = s;
            n = word.length();
            searchResult = false;
            isVisited = new int[n];
            backtrack(0, root);
            return searchResult;
        }
    }

   public boolean wordBreak(String s, List<String> wordDict) {
        Trie trie = new Trie();
        for (String word : wordDict) {
            trie.insert(word);
        }

        return trie.search(s);
    }
}
// @lc code=end

