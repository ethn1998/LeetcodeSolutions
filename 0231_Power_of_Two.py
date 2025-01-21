class Solution(object):
    def isPowerOfTwo(self, n):
        """
        :type n: int
        :rtype: bool
        """
        a = 1
        while a <= n:
            if a == n:
                return True
            else:
                a *= 2
        return False
