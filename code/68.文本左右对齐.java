import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=68 lang=java
 *
 * [68] 文本左右对齐
 *
 * https://leetcode-cn.com/problems/text-justification/description/
 *
 * algorithms
 * Hard (41.17%)
 * Likes:    34
 * Dislikes: 0
 * Total Accepted:    3.7K
 * Total Submissions: 8.9K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 * 
 * 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
 * 
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 * 
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 * 
 * 说明:
 * 
 * 
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 * 
 * 
 * 示例:
 * 
 * 输入:
 * words = ["This", "is", "an", "example", "of", "text", "justification."]
 * maxWidth = 16
 * 输出:
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入:
 * words = ["What","must","be","acknowledgment","shall","be"]
 * maxWidth = 16
 * 输出:
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。       
 * ⁠    第二行同样为左对齐，这是因为这行只包含一个单词。
 * 
 * 
 * 示例 3:
 * 
 * 输入:
 * words =
 * ["Science","is","what","we","understand","well","enough","to","explain",
 * "to","a","computer.","Art","is","everything","else","we","do"]
 * maxWidth = 20
 * 输出:
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 * 
 * 
 */

// @lc code=start
// class Solution {
//     private static String getSpace(int num) {
//         StringBuilder sb = new StringBuilder();
//         int i = 0;
//         while (i++ < num) {
//             sb.append(' ');
//         }
//         return sb.toString();
//     }

//     private static String getString(String[] words, int count, int start, int end, int maxWidth, boolean isLast) {
//         StringBuilder sb = new StringBuilder();

//         if (count > maxWidth) {
//             while (start < end) {
//                 sb.append(words[start++]);
//                 sb.append(' ');
//             }
//             sb.append(words[start++]);
//         } else {

//             if (isLast) {
//                 count = maxWidth;
//                 while (start < end) {
//                     count -= words[start].length() + 1;
//                     sb.append(words[start++]);
//                     sb.append(' ');
//                 }
//                 count -= words[start].length();

//                 sb.append(words[start++]);
//                 sb.append(getSpace(count));
//             } else {
//                 int space = 0, leftover = 0, r = 0, spaceInBet = 0;
//                 space = end - start + 1;
//                 count -= space;

//                 leftover = maxWidth - count;
//                 // but when we split space, no include last
//                 spaceInBet = (end == start) ? leftover : leftover / (end - start);
//                 // remainder
//                 r = (end == start) ? 0 : leftover % (end - start);

//                 while (start < end) {
//                     int numOfSpace = (r-- <= 0) ? spaceInBet : spaceInBet + 1;
//                     leftover -= words[start].length() + numOfSpace;
//                     sb.append(words[start++]);
//                     sb.append(getSpace(numOfSpace));
//                 }
//                 sb.append(words[start++]);
//                 if (leftover > 0) {
//                     sb.append(getSpace(leftover));
//                 }
//             }
//         }
//         return sb.toString();
//     }

//     public List<String> fullJustify(String[] words, int maxWidth) {
//         List<String> ans = new ArrayList<>();

//         int start = 0, end = 0, i = 0, len = words.length;

//         int count = 0;

//         while (i < len) {
//             int temp = count + words[i].length();
//             // false);
//             if (temp <= maxWidth) {
//                 count += words[i].length() + 1;
//                 i++;
//             } else {
//                 end = i - 1;

//                 if (end == len - 1)
//                     break;
//                 ans.add(getString(words, count, start, end, maxWidth, false));
//                 start = i;
//                 count = 0;
//             }

//         }
//         end = i - 1;
//         ans.add(getString(words, count, start, end, maxWidth, true));
//         return ans;
//     }
// }
class Solution {
    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> ans = new ArrayList<String>();
        int len = words.length;
        int start = 0;

        while (start < len) {
            int count = words[start].length();
            int end = start + 1;

            while (end < len) {
                if (words[end].length() + count + 1 > maxWidth)
                    break;
                count += words[end].length() + 1;
                end++;
            }

            StringBuilder sb = new StringBuilder();
            int diff = end - 1 - start;
            if (end == len || diff == 0) {
                for (int i = start; i < end; i++) {
                    sb.append(words[i] + ' ');
                }

                sb.deleteCharAt(sb.length() - 1);
                for (int i = sb.length(); i < maxWidth; i++) {
                    sb.append(' ');
                }
            } else {
                int leftover = maxWidth - count + diff;
                int spaces = leftover / diff;
                int r = leftover % diff;

                while (start < end - 1) {
                    int numOfSpace = (r-- <= 0) ? spaces : spaces + 1;
                    sb.append(words[start++]);
                    for (int i = 0; i < numOfSpace; i++) {
                        sb.append(' ');
                    }
                }
                sb.append(words[start++]);
            }
            ans.add(sb.toString());
            start = end;
        }
        return ans;
    }
}
// @lc code=end
