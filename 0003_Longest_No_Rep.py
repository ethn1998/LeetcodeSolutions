class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        if len(s) == 0:
            return 0
        if len(s) == 1:
            return 1
        fout = 0
        i = 0 #Left pointer
        j = 0 #Right pointer
        #l = 0 #Length of substring 
        seen = {}

        while j < len(s):
            char = s[j]
            if char in seen and seen[char] >= i: #Check if key exists
                if j-i > fout:
                    fout = j-i
                i = seen[char] + 1
            else:
                fout = max(fout,j-i+1)
            seen[char] = j
            j += 1
        return fout
        
