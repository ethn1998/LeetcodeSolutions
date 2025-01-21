class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = set(range(len(nums)+1)) #Expect to see everything
        for n in nums:
            s.remove(n) #Eliminate seen stuff
        return s.pop()
        
