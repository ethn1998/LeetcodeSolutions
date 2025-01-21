class Solution(object):
    def maximumGap(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) < 2:
            return 0
        nums.sort()
        delta = 0
        for i in list(range(len(nums)-1)):
            d = nums[i+1] - nums[i]
            if d > delta:
                delta = d
        return delta
        
