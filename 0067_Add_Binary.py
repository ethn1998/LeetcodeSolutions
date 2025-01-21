class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        fout = ""
        i = len(a)-1
        j = len(b)-1
        k = 0 #Carry forward
        while i >= 0 or j >= 0:
            v = k
            if i >= 0 :
                if a[i] == "1":
                    v += 1
                i -= 1                    
            if j >= 0:
                if b[j] == "1":
                    v += 1
                j -= 1

            d, k = v%2, v//2
            fout = "{}".format(d) + fout
        if k == 1:
            fout = "1" + fout
        return fout
