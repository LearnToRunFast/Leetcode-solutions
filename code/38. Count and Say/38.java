class Solution {
    public int maxSubArray(int[] nums) {
        if (nums == null || nums.length == 0) return 0;
        
        if (nums.length == 1) return nums[0];
        
        int max_so_far = nums[0], curr_max = nums[0];
        
        for (int i = 1; i < nums.length; i++) {
            // curr_max + nums[i] > nums[i]
            if (curr_max > 0) {
                curr_max += nums[i]; 
            } else {
                curr_max = nums[i];
            }
            max_so_far = Math.max(curr_max, max_so_far);
        }
        return max_so_far;
    }
}

