class Solution {

    public int searchInsert(int[] nums, int target) {
        
        int l = 0, h = nums.length - 1;
        int mid = 0;
        while (l <= h) {
            // what if l + h bigger than INT_MAX
            //mid = (l + h) / 2;
            //mid = l+ (h - l) / 2;
            mid = (l + h + 1) >>> 1;
            if (nums[mid] == target) return mid;
            if (nums[mid] > target) h = mid - 1;
            if (nums[mid] < target) l = mid + 1;
        }

        // 理由是对于 [1,3,5,6]，target = 2，返回大于等于 target 的第 1 个数的索引，此时应该返回 1
        // 在上面的 while (left <= right) 退出循环以后，right < left，right = 0 ，left = 1
        // 根据题意应该返回 left，
        // 如果题目要求你返回小于等于 target 的所有数里最大的那个索引值，应该返回 right


        return l;
    }
}