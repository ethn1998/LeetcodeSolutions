class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        #return x**n #?
        fout = 1
        isNegative = (n < 0)
        i = 0
        if isNegative:
            while i > n:
                fout /= x
                i -= 1
        else:
            while i < n:
                fout *= x
                i += 1
        return fout


        
