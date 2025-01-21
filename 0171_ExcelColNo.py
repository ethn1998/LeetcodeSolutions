class Solution(object):
    def titleToNumber(self, columnTitle):
        """
        :type columnTitle: str
        :rtype: int
        """
        columnTitle = columnTitle[::-1]
        fout = 0
        b = 1
        for char in columnTitle:
            fout += (ord(char) - 64)*b
            b *= 26
        return fout
        
