import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=127 lang=java
 *
 * [127] 单词接龙
 *
 * https://leetcode-cn.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (43.35%)
 * Likes:    629
 * Dislikes: 0
 * Total Accepted:    85.1K
 * Total Submissions: 187.8K
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * 给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord
 * 的最短转换序列的长度。转换需遵循如下规则：
 * 
 * 
 * 每次转换只能改变一个字母。
 * 转换过程中的中间单词必须是字典中的单词。
 * 
 * 
 * 说明:
 * 
 * 
 * 如果不存在这样的转换序列，返回 0。
 * 所有单词具有相同的长度。
 * 所有单词只由小写字母组成。
 * 字典中不存在重复的单词。
 * 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
 * 
 * 
 * 示例 1:
 * 
 * 输入:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 * 
 * 输出: 5
 * 
 * 解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
 * ⁠    返回它的长度 5。
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 * 
 * 输出: 0
 * 
 * 解释: endWord "cog" 不在字典中，所以无法进行转换。
 * 
 */

// @lc code=start
public class Solution {

    public int ladderLength(String beginWord, String endWord, List<String> wordList) {

        Set<String> wordSet = new HashSet<>(wordList);
        if (wordSet.size() == 0 || !wordSet.contains(endWord)) {
            return 0;
        }

        Set<String> visited = new HashSet<>();

        Set<String> beginSet = new HashSet<>();
        beginSet.add(beginWord);
        Set<String> endSet = new HashSet<>();
        endSet.add(endWord);
        
        int step = 0;
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            step++;

            if (beginSet.size() > endSet.size()) {
                Set<String> temp = beginSet;
                beginSet = endSet;
                endSet = temp;
            }

            // new set for next level, unlike queue, modify set during loop will lead to
            // unknown result.
            Set<String> newBeginSet = new HashSet<>();
            for (String currWord : beginSet) {
                if (isBeginSetMeetEndSet(currWord, newBeginSet, endSet, visited, wordSet)) {
                    return ++step;
                }
            }
            beginSet = newBeginSet;

        }
        return 0;
    }

    private boolean isBeginSetMeetEndSet(String currWord,
                                         Set<String> newBeginSet,
                                         Set<String> endSet, 
                                         Set<String> visited, 
                                         Set<String> wordSet) {

        char[] charArray = currWord.toCharArray();
        for (int i = 0; i < currWord.length(); i++) {
            char originChar = charArray[i];

            for (char c = 'a'; c <= 'z'; c++) {
                if (c == originChar)  continue;

                charArray[i] = c;
                String newWord = String.valueOf(charArray);
                if (wordSet.contains(newWord)) {
                    if (endSet.contains(newWord)) return true;

                    if (!visited.contains(newWord)) {
                        newBeginSet.add(newWord);
                        visited.add(newWord);
                    }
                }
            }
            // restore back original char
            charArray[i] = originChar;
        }
        return false;
    }
}


// @lc code=end

