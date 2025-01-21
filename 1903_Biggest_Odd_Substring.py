class Solution(object):
    def largestOddNumber(self, num):
        """
        :type num: str
        :rtype: str
        """
        i = len(num)-1
        while i >= 0:
            v = ord(num[i]) - 48
            if v % 2 != 0:
                return num[0:i+1]
            i -= 1
        return ""
        
