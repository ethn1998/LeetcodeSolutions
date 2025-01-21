class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        #Cut up magazine first
        cuts = {}
        for char in magazine:
            if char in cuts:
                cuts[char] += 1
            else:
                cuts[char] = 1
        for char in ransomNote:
            if cuts.get(char) is None:
                return False
            elif cuts[char] > 0:
                cuts[char] -= 1
            else:
                return False
        return True
