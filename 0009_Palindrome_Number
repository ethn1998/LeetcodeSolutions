import math

class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        #s = str(x)
        #i = 0
        #j = len(s)-1
        #while i <= j:
        #    if s[i] != s[j]:
        #        return False
        #    i += 1
        #    j -= 1
        #return True

        #Without converting integer to string
        if x < 0: #Common sense cases
            return False
        elif x == 0: #Prevents breaking use of log later on
            return True

        Omega = math.floor(math.log(x,10)) #Find magnitude of x
        i = 0
        j = Omega
        while i <= j:
            a = (x//(10**j))%10 #Left          
            b = (x//(10**i))%10 #Right
            if a != b:
                return False
            i += 1
            j -= 1
        return True


        
        
