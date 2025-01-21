class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        seen = set()
        seenagain = set()
        for n in nums:
            if n not in seen:
                seen.add(n)
            else:
                seenagain.add(n)
        fout = seen.difference(seenagain)
        return fout.pop()
        
