class Solution(object):
    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        if len(s) <= 10:
            return []
        fout = []
        repeats = set()
        mem = set()
        for i in range(len(s)-9):
            word = s[i:i+10]
            if word in mem:
                repeats.add(word)
            else:
                mem.add(word)
        for word in repeats:
            fout.append(word)
        return fout
