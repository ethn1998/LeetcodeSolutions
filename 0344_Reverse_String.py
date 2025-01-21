class Solution(object):
    def reverseString(self, s):
        """
        :type s: List[str]
        :rtype: None Do not return anything, modify s in-place instead.
        """
        i = 0
        while i < len(s)//2:
            v = s[i]
            s[i] = s[len(s)-1-i]
            s[len(s)-1-i] = v
            i += 1
