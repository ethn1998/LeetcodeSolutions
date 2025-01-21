class Solution(object):
    def getSneakyNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        seen = set()
        fout = []
        for n in nums:
            if n in seen:
                fout.append(n)
            else:
                seen.add(n)
        return fout
        
