class Solution(object):
    def convertToBase7(self, num):
        """
        :type num: int
        :rtype: str
        """
        if num == 0:
            return "0"
        isNegative = (num < 0)
        mag = abs(num)
        fout = ""
        while mag > 0:
            r = str(mag%7)
            fout = r + fout
            mag //= 7

        if isNegative:
            fout = "-" + fout
        return fout
