class Solution(object):
    def minMaxGame(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        while len(nums) > 1:
            newNums = []
            j = 0 #Index in next array
            while j < len(nums)//2:
                if j % 2 == 0:
                    newNums.append(min(nums[2*j],nums[2*j+1]))
                else:
                    newNums.append(max(nums[2*j],nums[2*j+1]))
                j += 1
            nums = newNums
        return nums[0]
        
