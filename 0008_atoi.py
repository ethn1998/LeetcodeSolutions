import math

asciivals = {43,45,48,49,50,51,52,53,54,55,56,57}
class Solution(object):
    def myAtoi(self, s):
        """
        :type s: str
        :rtype: int
        """
        isNegative = False #Sign
        m = 0 #Magnitude
        fout = 0
        #Whitespace
        s = s.lstrip() #Remove leading whitespace
        if s == "": #Degenerate case
            return 0

        #Signedness; check first character
        i0 = 0
        isNegative = False
        s0 = ord(s[0])
        if s0 not in asciivals:
            return 0
        elif s0 == 45:
            isNegative = True
            i0 = 1
        elif s0 == 43:
            i0 = 1 
        #Conversion
        for char in s[i0:]:
            if char.isdigit():
                m = 10*m + ord(char)-48
            else:
                break
        #Rounding?!
        if isNegative:
            if m > 2**31:
                return -2**31
            else:
                return -m
        else:
            if m >= 2**31:
                return 2**31-1
            else:
                return m
        
