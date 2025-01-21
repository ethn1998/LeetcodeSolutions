class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        Omega = set(range(1,len(nums)+1))
        fout = []
        for n in nums:
            Omega.discard(n)
        for v in Omega:
            fout.append(v)
        return fout
        
        
