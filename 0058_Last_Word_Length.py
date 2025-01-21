class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        s = s.strip()
        tlist = s.split(" ")
        return(len(tlist[-1]))


        
