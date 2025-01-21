class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """

        #Initialization
        fout = ""
        t = 0 #Ticker to alternate between words
        i = 0 #Index on first word
        j = 0 #Index on second word
        foutlength = len(word1 + word2) #Brute appending to check output length            

        while t < foutlength:
            if t%2 == 0:
                fout = fout + word1[i]
                i += 1
                if i > len(word1)-1:
                    return fout + word2[j:]
            else:
                fout = fout + word2[j]
                j += 1
                if j > len(word2)-1:
                    return fout + word1[i:]
            t = t+1
            
