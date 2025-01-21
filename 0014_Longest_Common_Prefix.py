class Solution(object):
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        fout = ""
        s = ""
        M = 201
        for word in strs: #Find and use shortest word
            if len(word)< M:
                s = word
                M = len(word)
        i = 0
        mybool = True #Boolean to check if we want to continue
        while i < M and mybool:
            char = s[i]
            for word in strs:
                if char != word[i]:
                    mybool = False
            if mybool:
                fout += char
                i += 1
            else:
                break
        return fout
        
