class Solution(object):
    def tribonacci(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n == 0:
            return 0
        if n == 1:
            return 1
        m  = 2
        t0 = 0
        t1 = 1
        t2 = 1
        while m < n:
            t0,t1,t2 = t1,t2,t0+t1+t2
            m += 1
        return t2
        
