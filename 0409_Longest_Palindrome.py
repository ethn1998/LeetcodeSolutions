class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: int
        """

        #Store frequency of occurrence of each letter
        freqs = {} #Use a dictionary
        for char in s:
            if char in freqs:
                freqs[char] +=1
            else:
                freqs[char] = 1

        O = False # Boolean to check if odds are present
        fout = 0
        for key in freqs.keys():
            m = freqs[key]
            if m % 2 == 0:
                fout += m
            else: #If odds are present
                O = True
                fout += m-1
        if O:
            return fout + 1
        else:
            return fout

        return O + E


        
