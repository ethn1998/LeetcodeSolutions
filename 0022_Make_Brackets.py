class Solution(object):
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        if n == 0:
            return [""]
        i = 1
        fout = ["()"]
        while i < n:
            #Generate new list
            bset = set()
            newfout = [] #Stores new extended strings
            for s in fout:
                for j in range(len(s)):
                    newb = s[0:j] + "()" + s[j:len(s)]
                    bset.add(newb)
            for s in bset:
                newfout.append(s) 
            fout = newfout
            i += 1

        return fout


        
