class Solution(object):
    def convertToTitle(self, columnNumber): 
        """
        :type columnNumber: int
        :rtype: str
        """
        #columnNumber = columnNumber - 1
        fout = ""
        while columnNumber > 0: #Excel columns basically use a base 26 system
            if columnNumber % 26 == 0:
                fout = "Z" + fout
                columnNumber = columnNumber//26 -1
            else:
                char = chr((columnNumber) % 26 + 64)
                fout = char + fout
                columnNumber //= 26
        return fout
