import java.util.List;
import java.util.Queue;

/*
 * @lc app=leetcode.cn id=17 lang=java
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (50.96%)
 * Likes:    453
 * Dislikes: 0
 * Total Accepted:    47.5K
 * Total Submissions: 93.3K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 * 
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 * 
 * 
 * 
 * 示例:
 * 
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 * 
 * 
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 * 
 */

// @lc code=start
// queue version
// class Solution {
//     public List<String> letterCombinations(String digits) {
//         String digitletter[] = {"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"};
//         List<String> res = new ArrayList<>();

//         if (digits.length() == 0) return res;   

//         Queue<String> queue = new LinkedList<String>();

//         queue.add("");
//         for(char digit: digits.toCharArray()) {
//             int len = queue.size();
//             for(int j = 0; j < len; j++) {
//                 String temp = queue.poll();

//                 for(char c: digitletter[digit - '0'].toCharArray()) {
//                     queue.add(temp + c);
//                 }
//             }
//         }

//         return new ArrayList<String> (queue);

//     }
// }

// backtracking version
class Solution {
    public List<String> letterCombinations(String digits) {
        List<String> res = new ArrayList<>();
        if (digits.length() == 0) return res;

        Map<Character, String> map = new HashMap<>() {
             {
                put('2', "abc");
                put('3', "def");
                put('4', "ghi");
                put('5', "jkl");
                put('6', "mno");
                put('7', "pqrs");
                put('8', "tuv");
                put('9', "wxyz");
             }
        };

        backtrack(res, map, digits, 0, new StringBuilder());

        return res;
    }

    private void backtrack(List<String> res, Map<Character, String> map, String digits, int index, StringBuilder builder) {
        if (index == digits.length()) {
            res.add(builder.toString());
            return;
        } 

        String letters = map.get(digits.charAt(index));
        for(char c : letters.toCharArray()) {
            builder.append(c);
            backtrack(res, map, digits, index + 1, builder);
            builder.deleteCharAt(index);
        }
        
    }
}

// @lc code=end

