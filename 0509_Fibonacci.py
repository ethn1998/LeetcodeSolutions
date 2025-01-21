class Solution(object):
    def fib(self, n):
        """
        :type n: int
        :rtype: int
        """
        f0 = 0 #n-1
        f1 = 1 #n
        if n == 0:
            return 0
        elif n == 1:
            return 1
        else:
            p = n-2
            i = 0
            while i <= p:
                f0,f1 = f1, f0+f1
                i += 1
            return f1 
        



        
