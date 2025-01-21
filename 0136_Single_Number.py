class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        s = set()
        for n in nums:
            if n in s:
                s.remove(n)
            else:
                s.add(n)
        return s.pop()

        
