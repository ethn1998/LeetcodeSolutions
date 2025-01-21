class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        fout = 1
        seen = set()
        for n in nums:
            seen.add(n)
            if n == fout:
                CHECKSEEN = True
                while CHECKSEEN:
                    if fout in seen:
                        seen.remove(fout)
                        fout += 1
                    else:
                        CHECKSEEN = False
        return fout
        
