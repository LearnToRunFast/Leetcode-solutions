class Solution {
    public int[] productExceptSelf(int[] nums) {
        if (nums == null || nums.length == 0) return null;
        int[] ans = new int[nums.length];
        
        for (int i = 0; i < nums.length; i++) {
            ans[i] = 1;
        }
        
        int temp = 1;
        for (int i = 0; i < nums.length; i++) {
            ans[i] = temp;
            temp *= nums[i];
        }
        temp = 1;
        for (int i = nums.length - 1; i >= 0; i--) {
            ans[i] *= temp;
            temp *= nums[i];

        }
        return ans;
    }
}