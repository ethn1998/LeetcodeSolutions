class Solution(object):
    def rotate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        # NOTE: Strategy inspired by Leetcode 969 Pancake Sort. Not the most efficient method, but a fun one!
        k = k % len(nums) # NOTE: Rotating by k is equivalent to not rotating. This prevents overflowing the array.
        # STEP 1: Pancake flip entire array. i.e. Reverse array
        i = 0
        j = len(nums)-1
        while i < j:
            nums[i],nums[j] = nums[j],nums[i]
            i += 1
            j -= 1
        # STEP 2: Pancake flip first slice
        i = 0 
        j = k-1
        while i < j:
            nums[i],nums[j] = nums[j],nums[i]
            i += 1
            j -= 1
        # STEP 3: Pancake flip second slice
        i = k 
        j = len(nums)-1
        while i < j:
            nums[i],nums[j] = nums[j],nums[i]
            i += 1
            j -= 1
