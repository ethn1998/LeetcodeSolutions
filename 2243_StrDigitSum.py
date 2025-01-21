class Solution(object):
    def digitSum(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """
        while len(s) > k:
            s1 = ""
            i = 0
            while i < len(s):
                v = 0
                subs = s[i:min(i+k,len(s))]
                for char in subs:
                    v += ord(char) - 48
                s1 += str(v)
                i += k
            s = s1
        return s
        
