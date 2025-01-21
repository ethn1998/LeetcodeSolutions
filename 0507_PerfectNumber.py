class Solution(object):
    def checkPerfectNumber(self, num):
        """
        :type num: int
        :rtype: bool
        """
        if num == 1: #Degenerate case
            return False
        s = 0
        p = 1
        n = num
        while p <= n:
            if num % p == 0:
                s += p + n
                if p == n or n == num:
                    s -= n
                n = num/p
            p += 1
        return s == num

        
