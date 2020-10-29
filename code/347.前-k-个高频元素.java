import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;

/*
 * @lc app=leetcode.cn id=347 lang=java
 *
 * [347] 前 K 个高频元素
 *
 * https://leetcode-cn.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (60.54%)
 * Likes:    553
 * Dislikes: 0
 * Total Accepted:    112.9K
 * Total Submissions: 182.9K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [1,1,1,2,2,3], k = 2
 * 输出: [1,2]
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [1], k = 1
 * 输出: [1]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
 * 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
 * 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
 * 你可以按任意顺序返回答案。
 * 
 * 
 */

// @lc code=start

// bucket sort
// class Solution {
//     public int[] topKFrequent(int[] nums, int k) {
//         Map<Integer, Integer> map = new HashMap<>();

//         for (int num : nums) {
//             map.put(num, map.getOrDefault(num, 0) + 1);
//         }
        
//         List<Integer>[] tmp = new List[nums.length + 1];

//         for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
//             int key = entry.getKey();
//             int value = entry.getValue();

//             if(tmp[value] == null){
//                 tmp[value] = new ArrayList<>();
//              } 


//             tmp[value].add(key);
//         }

//         int[] ans = new int[k];
//         k--; // off set the index

//         for (int i = tmp.length - 1; i >= 0 && k >= 0; i--) {
//             if (tmp[i] == null) continue;

//             for (int j = 0; j < tmp[i].size() && k >= 0; j++, k--) {
                
//                 ans[k] = tmp[i].get(j);
//             }
//         }

//         return ans;
//     }
// }

// hashmap + min heap
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }

        PriorityQueue<int[]> queue = new PriorityQueue<>(
            (int[] x, int[] y) -> {
                return x[1] - y[1];
            }
        );

        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            int key = entry.getKey();
            int value = entry.getValue();

            if (queue.size() == k) {
                if (queue.peek()[1] < value) {
                    queue.poll();
                    queue.add(new int[]{key, value});
                }
                continue;
            }
            queue.add(new int[]{key, value});
        }
        int[] ans = new int[k];

        for (int i = 0; i < k; i++) {
            ans[i] = queue.poll()[0];
        }
        return ans;
    }
}
// @lc code=end

