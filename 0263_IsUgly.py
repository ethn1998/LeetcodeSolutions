class Solution(object):
    def isUgly(self, n):
        """
        :type n: int
        :rtype: bool
        """
        #This is just a variation of fizzbuzz
        if n <= 0: #Constrain to positive integers
            return False
        for p in [2,3,5]:
            while n % p == 0: #While divisible
                n /= p
        return n == 1
