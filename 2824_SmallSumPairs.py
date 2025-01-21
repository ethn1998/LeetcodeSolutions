class Solution(object):
    def countPairs(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        fout = 0
        i = 0
        while i < len(nums):
            j = i+1
            while j < len(nums):
                s = nums[i] + nums[j]
                if s < target:
                    fout += 1               
                j += 1
            i += 1
        return fout
    
