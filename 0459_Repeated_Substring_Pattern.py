class Solution(object):
    def repeatedSubstringPattern(self, s):
        """
        :type s: str
        :rtype: bool
        """
        substrlen = 1
        while substrlen <= len(s)/2:
            substr = s[0:substrlen]
            genstr = ""
            while len(genstr) < len(s):
                genstr += substr
            if genstr == s:
                return True
            substrlen += 1
        return False
        
