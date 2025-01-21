# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num):

class Solution(object):
    def guessNumber(self, n):
        """
        :type n: int
        :rtype: int
        """
        i = 1
        j = n
        s = set() #Check seen numbers
        while i != j:
            k = (i+j)//2
            if k in s:
                k += 1                
            fb = guess(k)
            if fb < 0:
                j = k
            elif fb > 0:
                i = k
            else:
                return k
            s.add(k)
        return i
