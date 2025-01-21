class Solution(object):
    def percentageLetter(self, s, letter):
        """
        :type s: str
        :type letter: str
        :rtype: int
        """
        count = 0
        for char in s:
            if char == letter:
                count += 1
        return 100 * count //len(s)
        
