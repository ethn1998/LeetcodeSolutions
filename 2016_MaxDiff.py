class Solution(object):
    def maximumDifference(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        fout = -1
        minSeen = nums[0]
        i,j = 0,1
        while j < len(nums):
            if nums[j] < minSeen:
                minSeen = nums[j]
                i = j
                #j += 1
            else:
                diff = nums[j] - nums[i]
                if diff > fout and diff != 0:
                    fout = nums[j] - nums[i]
            j += 1
        return fout


        
