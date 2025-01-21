class Solution(object):
    def minimumLength(self, s):
        """
        :type s: str
        :rtype: int
        """
        seen = {} #Dictionary to store seen characters and their indices.
        for char in s:
            if seen.get(char) != None:
                seen[char] += 1
            else:
                seen[char] = 1
        fout = 0
        for char in seen:
            if seen[char]%2 != 0:
                fout += 1
            else:
                fout += 2
        return fout

