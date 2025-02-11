class Solution(object):
    def removeOccurrences(self, s, part):
        """
        :type s: str
        :type part: str
        :rtype: str
        """
        fout = ""
        i = 0
        for i in range(len(s)):
            fout += s[i]
            if len(fout) >= len(part):
                substr = fout[-len(part):]
                if substr == part:
                    fout = fout[:-len(part)]
        return fout
