class Solution(object):
    def findContentChildren(self, g, s):
        """
        :type g: List[int]
        :type s: List[int]
        :rtype: int
        """
        fout = 0 #Initialize output
        g.sort() #Clean childgreed list by sorting
        s.sort() #Clean cookiesize list by sorting
        # NOTE: Implement two pointer solution: One pointing to cookiesize list, the other pointing to childgreedlist.
        i = 0 #Cookie size 
        j = 0 #Child greed
        while i < len(s) and j < len(g): #We only have at most len(s) cookies and at most len(g) kids.
            if s[i] >= g[j]: #If jth child accepts cookie
                fout += 1 #Feed ith child jth cookie
                j += 1 #Try to feed next child
            i += 1 #Go to bigger cookie if child eats or rejects cookie
        return fout
            
