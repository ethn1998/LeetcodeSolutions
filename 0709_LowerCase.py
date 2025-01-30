class Solution(object):
    def toLowerCase(self, s):
        """
        :type s: str
        :rtype: str
        """
        #NOTE: Challenge: do not use any built in string functions
        i = 0
        while i < len(s):
            if ord(s[i]) >= 65 and ord(s[i]) <= 90:
                s = "{}{}{}".format(s[:i],chr(ord(s[i]) + 32),s[i+1:])
            i += 1
        return s
