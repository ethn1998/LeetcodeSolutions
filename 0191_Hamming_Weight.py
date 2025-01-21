class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        fout = 0
        while n>0:
            q = n%2
            if q == 1:
                fout += 1
            n //= 2

        return fout
