class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        a = 1
        while a <= n:
            if a == n:
                return True
            else:
                a *= 3
        return False
