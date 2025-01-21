import math

class Solution(object):
    def addDigits(self, num):
        """
        :type num: int
        :rtype: int
        """
        if num == 0: #Degenerate case
            return 0
        while True: #We want to do this endlessly
            p = math.log10(num)
            i = 0
            s = 0
            while i <= p:
                s += (num//10**i) % 10
                i += 1
            if s < 10:
                return s
            num = s                

        
