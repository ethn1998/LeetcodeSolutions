class Solution(object):
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        """
        NOTE: sort and count pairwise?
        """
        nums.sort()
        fout = 0
        i = 0
        while i < len(nums):
            fout += nums[i]
            i += 2
        return fout
        
