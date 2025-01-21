class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        #Sqrt(x) is monotone increasing in (x)
        u = 0
        while u*u < x:
            u += 1
        return u-1
