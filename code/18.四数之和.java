import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*
 * @lc app=leetcode.cn id=18 lang=java
 *
 * [18] 四数之和
 *
 * https://leetcode-cn.com/problems/4sum/description/
 *
 * algorithms
 * Medium (35.93%)
 * Likes:    294
 * Dislikes: 0
 * Total Accepted:    34.9K
 * Total Submissions: 96.9K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c
 * + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
 * 
 * 注意：
 * 
 * 答案中不可以包含重复的四元组。
 * 
 * 示例：
 * 
 * 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
 * 
 * 满足要求的四元组集合为：
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 * 
 * 
 */

// @lc code=start
class Solution {

    int len = -1;
    public List<List<Integer>> fourSum(int[] nums, int target) {
        len = nums.length;
        Arrays.sort(nums);
        return kSum(nums, target, 4, 0);
    }
    private ArrayList<List<Integer>> kSum(int[] nums, int target, int k, int head) {
        ArrayList<List<Integer>> ans = new ArrayList<>();
        // unexpected situation 
        if (head >= len) {
            return ans;
        }
        if (k == 2) {
            int i = head, j = len - 1;
            while(i < j) {
                if (target - nums[i] == nums[j]) {
                    List<Integer> temp = new ArrayList<>();
                    temp.add(nums[i]);
                    temp.add(nums[j]);
                    ans.add(temp);
                    //skip duplication
                    while(i < j && nums[i] == nums[i + 1]) i++;
                    while(i < j && nums[j] == nums[j - 1]) j--;
                    i++;
                    j--;

                } else if (target - nums[i] > nums[j]) {
                    i++;
                } else {
                    j--;
                }
            }
        } else {
            for(int i = head; i < len - k + 1; i++) {
                ArrayList<List<Integer>> temp = kSum(nums, target - nums[i], k - 1, i + 1);
                if (temp.size() > 0) {
                    for(List<Integer> sub : temp) {
                        sub.add(0, nums[i]);
                    }
                    ans.addAll(temp);
                }
                while(i < len - 1 && nums[i] == nums[i + 1]) {
                    i++;
                }
            }
        }

        return ans;
    }
}
// @lc code=end

