class Solution(object):
    def climbStairs(self, n):
        """
        :type n: int
        :rtype: int
        """
        #if n <= 1:
        #    return n
        m = 0

        x1 = 0
        x2 = 1
        while m<n:
            x1,x2 = x2, x1+x2
            m += 1
        return x2
