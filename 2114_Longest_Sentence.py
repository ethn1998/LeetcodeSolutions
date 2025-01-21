class Solution(object):
    def mostWordsFound(self, sentences):
        """
        :type sentences: List[str]
        :rtype: int
        """
        fout = 0
        for s in sentences:
            words = s.split(" ")
            if len(words) > fout:
                fout = len(words)
        return fout
