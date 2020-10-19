import java.util.Arrays;

/*
 * @lc app=leetcode.cn id=215 lang=java
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.36%)
 * Likes:    746
 * Dislikes: 0
 * Total Accepted:    213.2K
 * Total Submissions: 331K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 * 
 * 示例 1:
 * 
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 * 
 * 
 * 示例 2:
 * 
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 * 
 * 说明: 
 * 
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 * 
 */

// @lc code=start
class Solution {

    private void insertionSort(int[] nums, int l, int r) {
        for (int i = l + 1; i <= r; i++) {
            int key = nums[i];
            int j = i - 1;
            while (j >= l && key < nums[j]) {
                nums[j + 1] = nums[j];
                j--;
            }
            nums[j + 1] = key;
        }
    }
    private void swap(int[] nums, int i, int j) {
        int tmp = nums[i];
        nums[i] = nums[j];
        nums[j] = tmp;
    }

    private int medianOfMedians(int[] nums, int l, int r) {
        if (l == r) return nums[l];

        int replaceIdx = l;
        int firstIdx = replaceIdx;
        while (l + 5 < r) {
            insertionSort(nums, l, l + 4);
            swap(nums, replaceIdx, l + 2);
            replaceIdx++;   
            l += 5;
        }
        if (l == r) {
            swap(nums, replaceIdx, l);
        } else {
            insertionSort(nums, l, r);
            swap(nums, replaceIdx, l + (r - l) / 2);
        }

        return medianOfMedians(nums, firstIdx, replaceIdx);
    }

    private int partition(int[] nums, int l, int r, int pivot) {
        // for this version partition, we assume pivot locationed at l.
        while (l < r) {
            while (l < r && nums[r] >= pivot) {
                r--;
            }
            nums[l] = nums[r];

            while (l < r && nums[l] <= pivot) {
                l++;
            }
            nums[r] = nums[l];
        }
        nums[l] = pivot;
        return l;
    }

    private int bfptr(int[] nums, int k, int l, int r) {
        if (nums.length <= 5) {
            Arrays.sort(nums);
            return nums[k];
        }

        while (l < r) {
            int pivot = medianOfMedians(nums, l, r);
            //int pivot =nums[l];

            int mid = partition(nums, l, r, pivot);
            if (mid == k) return nums[mid];
            if (mid < k) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        return nums[l];
    }
    public int findKthLargest(int[] nums, int k) {
        return bfptr(nums, nums.length - k, 0, nums.length - 1);
    }
}
// @lc code=end

