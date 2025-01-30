class Solution(object):
    def countAndSay(self, n):
        """
        :type n: int
        :rtype: str
        """
        fout = "1"
        k = 1
        while k < n:
            newstr = ""
            i = 0
            j = 0
            while j < len(fout):
                if fout[i] != fout[j]:
                    newstr += "{}{}".format(j-i,fout[i])
                    i = j
                j += 1
            newstr += "{}{}".format(j-i,fout[i])
            fout = newstr
            k += 1
        return fout
        
