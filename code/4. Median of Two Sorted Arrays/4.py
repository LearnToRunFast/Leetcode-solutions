class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        # make sure first array is shorter than second one 
        if len(nums1) > len(nums2):
          nums1, nums2 = nums2, nums1
        
        m, n = len(nums1), len(nums2)
        
        left, right = 0, m
        
        left_total = (m + n + 1) >> 1

        while left < right:
          i = (left + right) >> 1
          j = left_total - i
          
          if nums2[j - 1] > nums1[i]:
            left = i + 1
          else:
            right = i
            
        i = left
        j = left_total - left
        
        #edge cases, when left side values all in one array.
        nums1_left_max = float('-inf') if i == 0 else nums1[i - 1]
        nums1_right_min = float('inf') if i == m else nums1[i]
        
        nums2_left_max = float('-inf') if j == 0 else nums2[j - 1]
        nums2_right_min = float('inf') if j == n else nums2[j]
        
        if (m + n) & 1:
          return max(nums1_left_max, nums2_left_max)
        else:
          return (max(nums1_left_max, nums2_left_max) + min(nums1_right_min, nums2_right_min)) / 2
          
          