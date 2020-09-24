class Solution {
    public int search(int[] nums, int target) {
        if (nums == null || nums.length == 0) return -1;
        int i = 0, j = nums.length - 1;
        while (i < j) {
            int mid = (i + j)  >>> 1;
            if (nums[mid] > nums[j]) {
                i = mid + 1;
            } else {
                j = mid;
            }
        }
        if (target == nums[i]) return i;
        
        if (target > nums[nums.length - 1]) {
            j = i;
            i = 0;

        } else if (target < nums[nums.length - 1]) {
            i = i;
            j = nums.length - 1;

        } else {
            return nums.length - 1;
        }

        while (i <= j) {
            int mid = (i + j) >>> 1;

            if (nums[mid] > target) {
                j = mid -1;
            } else if (nums[mid] < target) {
                i = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;
    }
}

class Solution {
    public int search(int[] nums, int target) {
        int i = 0, j = nums.length - 1;
        while (i <= j) {
            int mid = (i + j)  >>> 1;
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < nums[j]) {
                // right is ordered
                // nums[mid] -- target -- nums[j]
                if (nums[mid] < target && target <= nums[j]) {
                   i = mid + 1; 
                } else {
                    j = mid - 1;
                }
            } else {
                // left is ordered
                // nums[i] --- target --- nums[mid]
                if (nums[mid] > target && target >= nums[i]) {
                    j = mid - 1;
                } else {
                    i = mid + 1;
                }
            }
        }
        return -1;
    }
}