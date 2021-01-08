import java.util.Deque;
import java.util.Queue;
import java.util.Set;

/*
 * @lc app=leetcode.cn id=752 lang=java
 *
 * [752] 打开转盘锁
 *
 * https://leetcode-cn.com/problems/open-the-lock/description/
 *
 * algorithms
 * Medium (49.02%)
 * Likes:    208
 * Dislikes: 0
 * Total Accepted:    31.1K
 * Total Submissions: 63.3K
 * Testcase Example:  '["0201","0101","0102","1212","2002"]\n"0202"'
 *
 * 你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8',
 * '9' 。每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
 * 
 * 锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
 * 
 * 列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
 * 
 * 字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
 * 输出：6
 * 解释：
 * 可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
 * 注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
 * 因为当拨动到 "0102" 时这个锁就会被锁定。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: deadends = ["8888"], target = "0009"
 * 输出：1
 * 解释：
 * 把最后一位反向旋转一次即可 "0000" -> "0009"。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"],
 * target = "8888"
 * 输出：-1
 * 解释：
 * 无法旋转到目标数字且不被锁定。
 * 
 * 
 * 示例 4:
 * 
 * 
 * 输入: deadends = ["0000"], target = "8888"
 * 输出：-1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 死亡列表 deadends 的长度范围为 [1, 500]。
 * 目标数字 target 不会在 deadends 之中。
 * 每个 deadends 和 target 中的字符串的数字会在 10,000 个可能的情况 '0000' 到 '9999' 中产生。
 * 
 * 
 */

// @lc code=start
// class Solution {
//     private String increase(String item, int idx) {
//        StringBuilder sb = new StringBuilder(item);
//        sb.setCharAt(idx, sb.charAt(idx) == '9' ? '0' : (char) (sb.charAt(idx) + 1));
//        return sb.toString(); 
//     }
//     private String decrease(String item, int idx) {
//         StringBuilder sb = new StringBuilder(item);
//         sb.setCharAt(idx, sb.charAt(idx) == '0' ? '9' : (char) (sb.charAt(idx) - 1));
//         return sb.toString(); 
//      }
//     public int openLock(String[] deadends, String target) {
//         Set<String> visited = new HashSet<>();
//         for(String deadend : deadends) {
//             visited.add(deadend);
//         }
//         String start = "0000";
//         if (visited.contains(start)) {
//             return -1;
//         }
//         Queue<String> queue = new LinkedList<>();
//         queue.add(start);
//         visited.add(start);
//         int step = 0;
//         while (!queue.isEmpty()) {
//             int size = queue.size();
//             for (int i = 0; i < size; i++) {
//                 String curr = queue.poll();
                
//                 if (curr.equals(target)) return step;
//                 for (int j = 0; j < 4; j++) {
//                     String next = increase(curr, j);
//                     if (!visited.contains(next)) {
//                         queue.offer(next);
//                         visited.add(next); 
//                     }
//                     next = decrease(curr, j);
//                     if (!visited.contains(next)) {
//                         queue.offer(next);
//                         visited.add(next);
//                     }
//                 }
//             }
//             step++;
//         }
//         return -1;
//     }
// }
// two way bfs
class Solution {
    private String increase(String item, int idx) {
       StringBuilder sb = new StringBuilder(item);
       sb.setCharAt(idx, sb.charAt(idx) == '9' ? '0' : (char) (sb.charAt(idx) + 1));
       return sb.toString(); 
    }
    private String decrease(String item, int idx) {
        StringBuilder sb = new StringBuilder(item);
        sb.setCharAt(idx, sb.charAt(idx) == '0' ? '9' : (char) (sb.charAt(idx) - 1));
        return sb.toString(); 
     }
    public int openLock(String[] deadends, String target) {
        Set<String> visited = new HashSet<>();
        Set<String> q1 = new HashSet<>();
        Set<String> q2 = new HashSet<>();
        for(String deadend : deadends) {
            visited.add(deadend);
        }
        String start = "0000";
        if (visited.contains(start)) {
            return -1;
        }
        q1.add(start);
        q2.add(target);
        int step = 0;
        while (!q1.isEmpty() && !q2.isEmpty()) {
            Set<String> temp = new HashSet<>();
            for (String curr : q1) {
                if (q2.contains(curr)) return step;
                visited.add(curr);

                for (int j = 0; j < 4; j++) {
                    String next = increase(curr, j);
                    if (!visited.contains(next)) {
                        temp.add(next);
                    }
                    next = decrease(curr, j);
                    if (!visited.contains(next)) {
                        temp.add(next);
                    }
                }

            }

            step++;

            if (q2.size() > temp.size()) {
                q1 = temp;
            } else {
                q1 = q2;
                q2 = temp;
            }
        }
        return -1;
    }
}
// @lc code=end

