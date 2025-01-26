class Solution(object):
    def triangularSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        prevrow = nums
        i = len(nums)
        while i > 1:
            newrow = []
            for j in range(len(prevrow)-1):
                newrow.append((prevrow[j]+prevrow[j+1])%10)
            prevrow = newrow
            i -= 1
        return prevrow[-1]
