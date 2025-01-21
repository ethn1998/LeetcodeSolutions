class Solution(object):
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        d = {"I":1, "V":5, "X":10, "L":50, "C":100, "D":500, "M":1000}
        fout = 0
        u = 0
        v = 0
        #i = 0 #Index tracker

        while len(s)>1 :
            u = d.get(s[0])
            v = d.get(s[1])
            if u < v:
                fout += v-u
                s = s[2:] #jump
            else:
                fout += u
                s = s[1:]
        if len(s) == 1:
            fout += d.get(s) #This is if one character is left
        return fout


        
