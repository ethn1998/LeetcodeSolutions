class Solution(object):
    def isPowerOfFour(self, n):
        """
        :type n: int
        :rtype: bool
        """
        a = 1
        while a <= n:
            if a == n:
                return True
            else:
                a *= 4
        return False
        
