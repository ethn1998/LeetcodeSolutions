class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        i = 0 #Tracks where in string t are we.

        #Common sense cases
        if s == "":
            return True
        if len(s) > len(t):
            return False

        #General cases
        for char in s:
            k = t.find(char,i) #Try to get first occurence of letter l in t starting from i
            if k < 0: #If there are no more occurences
                return False
            i = k+1 #Update to start from the latest point
        return True
        
